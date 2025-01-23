#ifndef platform_ios_h
#define platform_ios_h

#import <Foundation/Foundation.h>
#import <NetworkExtension/NetworkExtension.h>
#import <SystemConfiguration/SystemConfiguration.h>

// Add version info
#define MIHOMO_VERSION "1.0.0"

// Expanded error types
typedef enum {
    TunnelErrorNone = 0,
    TunnelErrorPermissionDenied = 1,
    TunnelErrorNetworkError = 2,
    TunnelErrorSystemError = 3,
    TunnelErrorInvalidConfiguration = 4,
    TunnelErrorTunnelStartFailed = 5,
    TunnelErrorDNSConfigFailed = 6,
    TunnelErrorProxyConfigFailed = 7,
    TunnelErrorMemoryError = 8,
    TunnelErrorTimeout = 9,
    TunnelErrorAlreadyRunning = 10
} TunnelError;

// Function declarations with error handling
TunnelError initializeNetworkExtension(void);
TunnelError setupTunnel(const char* tunName);
TunnelError configureDNS(const char* servers);
TunnelError setSystemProxy(const char* host, int port);

// Improved TUN interface configuration
@interface TunnelConfiguration : NSObject
@property (nonatomic, strong) NSString *name;
@property (nonatomic, assign) int mtu;
@property (nonatomic, strong) NSArray<NSString *> *addresses;
@property (nonatomic, strong) NSArray<NSString *> *dnsServers;
@property (nonatomic, assign) BOOL includedRoutes;
@property (nonatomic, assign) BOOL excludedRoutes;
@end

// Enhanced Network Extension provider interface
@interface PacketTunnelProvider : NEPacketTunnelProvider
- (void)startTunnelWithConfiguration:(TunnelConfiguration *)config 
                  completionHandler:(void (^)(TunnelError error))completion;
- (void)stopTunnelWithCompletionHandler:(void (^)(TunnelError error))completion;
- (void)updateTunnelConfiguration:(TunnelConfiguration *)config;
@end

#endif /* platform_ios_h */ 