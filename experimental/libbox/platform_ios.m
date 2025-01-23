#import "platform_ios.h"

// Global tunnel provider instance
static PacketTunnelProvider *tunnelProvider = nil;

// Initialize NetworkExtension framework
TunnelError initializeNetworkExtension(void) {
    @autoreleasepool {
        if (@available(iOS 9.0, *)) {
            __block TunnelError result = TunnelErrorNone;
            dispatch_semaphore_t semaphore = dispatch_semaphore_create(0);
            
            [NEVPNManager loadAllFromPreferencesWithCompletionHandler:^(NSArray<NEVPNManager *> * _Nullable managers, NSError * _Nullable error) {
                if (error) {
                    switch (error.code) {
                        case NEVPNErrorConfigurationInvalid:
                            result = TunnelErrorInvalidConfiguration;
                            break;
                        case NEVPNErrorConfigurationDisabled:
                            result = TunnelErrorPermissionDenied;
                            break;
                        default:
                            result = TunnelErrorSystemError;
                    }
                }
                dispatch_semaphore_signal(semaphore);
            }];
            
            // Wait with timeout
            if (dispatch_semaphore_wait(semaphore, dispatch_time(DISPATCH_TIME_NOW, 5 * NSEC_PER_SEC)) != 0) {
                return TunnelErrorTimeout;
            }
            return result;
        } else {
            return TunnelErrorSystemError;
        }
    }
}

// Setup tunnel interface with enhanced error handling
TunnelError setupTunnel(const char* tunName) {
    @autoreleasepool {
        if (!tunName) {
            return TunnelErrorInvalidConfiguration;
        }
        
        NSString *name = [NSString stringWithUTF8String:tunName];
        if (!name) {
            return TunnelErrorMemoryError;
        }
        
        // Check if tunnel is already running
        if (tunnelProvider && tunnelProvider.isEnabled) {
            return TunnelErrorAlreadyRunning;
        }
        
        // Create tunnel configuration with validation
        TunnelConfiguration *config = [[TunnelConfiguration alloc] init];
        if (!config) {
            return TunnelErrorMemoryError;
        }
        
        config.name = name;
        config.mtu = 1500;
        
        __block TunnelError result = TunnelErrorNone;
        dispatch_semaphore_t semaphore = dispatch_semaphore_create(0);
        
        // Initialize and start tunnel with timeout
        if (!tunnelProvider) {
            tunnelProvider = [[PacketTunnelProvider alloc] init];
        }
        
        [tunnelProvider startTunnelWithConfiguration:config completionHandler:^(TunnelError error) {
            result = error;
            dispatch_semaphore_signal(semaphore);
        }];
        
        if (dispatch_semaphore_wait(semaphore, dispatch_time(DISPATCH_TIME_NOW, 10 * NSEC_PER_SEC)) != 0) {
            return TunnelErrorTimeout;
        }
        
        return result;
    }
}

// Configure DNS with enhanced error handling
TunnelError configureDNS(const char* servers) {
    @autoreleasepool {
        if (!servers) {
            return TunnelErrorInvalidConfiguration;
        }
        
        NSString *serversString = [NSString stringWithUTF8String:servers];
        if (!serversString) {
            return TunnelErrorMemoryError;
        }
        
        NSArray *dnsServers = [serversString componentsSeparatedByString:@","];
        if (!dnsServers || dnsServers.count == 0) {
            return TunnelErrorInvalidConfiguration;
        }
        
        // Validate DNS servers
        for (NSString *server in dnsServers) {
            if (![self isValidIPAddress:server]) {
                return TunnelErrorInvalidConfiguration;
            }
        }
        
        if (!tunnelProvider) {
            return TunnelErrorSystemError;
        }
        
        __block TunnelError result = TunnelErrorNone;
        dispatch_semaphore_t semaphore = dispatch_semaphore_create(0);
        
        TunnelConfiguration *config = [[TunnelConfiguration alloc] init];
        config.dnsServers = dnsServers;
        
        [tunnelProvider updateTunnelConfiguration:config withCompletionHandler:^(NSError *error) {
            if (error) {
                result = TunnelErrorDNSConfigFailed;
            }
            dispatch_semaphore_signal(semaphore);
        }];
        
        if (dispatch_semaphore_wait(semaphore, dispatch_time(DISPATCH_TIME_NOW, 5 * NSEC_PER_SEC)) != 0) {
            return TunnelErrorTimeout;
        }
        
        return result;
    }
}

// Helper method to validate IP addresses
- (BOOL)isValidIPAddress:(NSString *)ipAddress {
    struct in_addr sa;
    struct in6_addr sa6;
    
    if (inet_pton(AF_INET, [ipAddress UTF8String], &sa) == 1) {
        return YES;
    }
    if (inet_pton(AF_INET6, [ipAddress UTF8String], &sa6) == 1) {
        return YES;
    }
    return NO;
}

// Set system proxy
TunnelError setSystemProxy(const char* host, int port) {
    @autoreleasepool {
        NSString *proxyHost = [NSString stringWithUTF8String:host];
        
        NSDictionary *proxySettings = @{
            (NSString *)kCFProxyHostNameKey: proxyHost,
            (NSString *)kCFProxyPortNumberKey: @(port),
            (NSString *)kCFProxyTypeKey: (NSString *)kCFProxyTypeHTTP
        };
        
        CFDictionaryRef dictRef = (__bridge CFDictionaryRef)proxySettings;
        SCDynamicStoreRef store = SCDynamicStoreCreate(NULL, CFSTR("Mihomo"), NULL, NULL);
        
        if (store) {
            if (SCDynamicStoreSetValue(store, kSCPrefNetworkServices, dictRef)) {
                CFRelease(store);
                return TunnelErrorNone;
            }
            CFRelease(store);
        }
        
        return TunnelErrorSystemError;
    }
}

// PacketTunnelProvider implementation
@implementation PacketTunnelProvider

- (void)startTunnelWithConfiguration:(TunnelConfiguration *)config 
                  completionHandler:(void (^)(TunnelError))completion {
    NEPacketTunnelNetworkSettings *settings = [[NEPacketTunnelNetworkSettings alloc] initWithTunnelRemoteAddress:@"127.0.0.1"];
    
    // Configure interface
    settings.MTU = @(config.mtu);
    settings.tunnelOverheadBytes = @(0);
    
    // Configure DNS if provided
    if (config.dnsServers.count > 0) {
        NEDNSSettings *dnsSettings = [[NEDNSSettings alloc] initWithServers:config.dnsServers];
        settings.DNSSettings = dnsSettings;
    }
    
    // Apply settings
    [self setTunnelNetworkSettings:settings completionHandler:^(NSError * _Nullable error) {
        if (error) {
            completion(TunnelErrorNetworkError);
            return;
        }
        completion(TunnelErrorNone);
    }];
}

- (void)stopTunnelWithCompletionHandler:(void (^)(TunnelError))completion {
    [self setTunnelNetworkSettings:nil completionHandler:^(NSError * _Nullable error) {
        completion(error ? TunnelErrorSystemError : TunnelErrorNone);
    }];
}

- (void)updateTunnelConfiguration:(TunnelConfiguration *)config {
    // Update tunnel settings as needed
    NEPacketTunnelNetworkSettings *settings = [[NEPacketTunnelNetworkSettings alloc] initWithTunnelRemoteAddress:@"127.0.0.1"];
    
    if (config.dnsServers.count > 0) {
        NEDNSSettings *dnsSettings = [[NEDNSSettings alloc] initWithServers:config.dnsServers];
        settings.DNSSettings = dnsSettings;
    }
    
    [self setTunnelNetworkSettings:settings completionHandler:nil];
}

@end

@implementation TunnelConfiguration
@end 