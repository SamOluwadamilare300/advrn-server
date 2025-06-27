# Advrn API Endpoints Analysis & Implementation Plan

https://claude.ai/public/artifacts/b0416754-1e92-484a-aff0-0a1cbd2f61e9
https://claude.ai/public/artifacts/54a662c0-ba66-4acc-8ed6-93183e3065e7

## 📊 Current Endpoints Assessment

### ✅ **Well-Implemented Modules**
- **Authentication**: Complete with social login, password reset, token refresh
- **User Management**: Basic profile operations, saved properties, notifications
- **Property Management**: CRUD operations, search by location
- **Messaging**: Conversation and message handling
- **Reviews**: Property review system
- **Location Services**: Autocomplete and search functionality

### ⚠️ **Missing Critical Modules for MVP**

## 🚀 Priority 1: Essential Missing Endpoints

### 1. **Payment Module** (`/api/payment`)
```go
// Critical for marketplace functionality
payment.POST("/initialize", accessTokenVerifierMiddleware, routes.InitializePayment)
payment.POST("/verify", accessTokenVerifierMiddleware, routes.VerifyPayment)
payment.GET("/user/{id}/history", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserPaymentHistory)
payment.POST("/refund/{id}", accessTokenVerifierMiddleware, routes.ProcessRefund)
payment.GET("/escrow/{property_id}", accessTokenVerifierMiddleware, routes.GetEscrowStatus)
payment.POST("/recurring/setup", accessTokenVerifierMiddleware, routes.SetupRecurringPayment)
```

### 2. **Rental Applications Module** (`/api/applications`)
```go
// Core marketplace feature - currently missing
applications.POST("/", accessTokenVerifierMiddleware, routes.CreateApplication)
applications.GET("/property/{id}", accessTokenVerifierMiddleware, routes.GetApplicationsByPropertyID)
applications.GET("/user/{id}", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserApplications)
applications.PATCH("/{id}/status", accessTokenVerifierMiddleware, routes.UpdateApplicationStatus)
applications.POST("/{id}/documents", accessTokenVerifierMiddleware, routes.UploadApplicationDocuments)
```

### 3. **User Verification Module** (`/api/verification`)
```go
// Essential for trust and safety
verification.POST("/identity", accessTokenVerifierMiddleware, routes.SubmitIdentityVerification)
verification.POST("/employment", accessTokenVerifierMiddleware, routes.SubmitEmploymentVerification)
verification.POST("/landlord", accessTokenVerifierMiddleware, routes.SubmitLandlordVerification)
verification.POST("/employer", accessTokenVerifierMiddleware, routes.SubmitEmployerVerification)
verification.GET("/status/{user_id}", accessTokenVerifierMiddleware, routes.GetVerificationStatus)
verification.PATCH("/{id}/approve", accessTokenVerifierMiddleware, routes.ApproveVerification) // Admin only
```

### 4. **Lease Management Module** (`/api/lease`)
```go
// Critical for rental lifecycle
lease.POST("/", accessTokenVerifierMiddleware, routes.CreateLease)
lease.GET("/{id}", accessTokenVerifierMiddleware, routes.GetLeaseDetails)
lease.GET("/user/{id}", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserLeases)
lease.PATCH("/{id}/sign", accessTokenVerifierMiddleware, routes.SignLease)
lease.POST("/{id}/renew", accessTokenVerifierMiddleware, routes.RenewLease)
lease.PATCH("/{id}/terminate", accessTokenVerifierMiddleware, routes.TerminateLease)
```

## 🎯 Priority 2: Important Missing Endpoints

### 5. **Insurance Module** (`/api/insurance`)
```go
insurance.GET("/plans", routes.GetInsurancePlans)
insurance.POST("/quote", accessTokenVerifierMiddleware, routes.GetInsuranceQuote)
insurance.POST("/purchase", accessTokenVerifierMiddleware, routes.PurchaseInsurance)
insurance.GET("/user/{id}/policies", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserPolicies)
insurance.POST("/claim", accessTokenVerifierMiddleware, routes.SubmitClaim)
```

### 6. **Logistics Module** (`/api/logistics`)
```go
logistics.GET("/providers", routes.GetLogisticsProviders)
logistics.POST("/quote", accessTokenVerifierMiddleware, routes.GetMovingQuote)
logistics.POST("/book", accessTokenVerifierMiddleware, routes.BookMovingService)
logistics.GET("/user/{id}/bookings", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserBookings)
logistics.PATCH("/{id}/track", routes.TrackMovingService)
```

### 7. **Virtual Tours Module** (`/api/virtual-tours`)
```go
// For your virtual tour requirement
virtualTours.POST("/property/{id}", accessTokenVerifierMiddleware, routes.CreateVirtualTour)
virtualTours.GET("/property/{id}", routes.GetVirtualTour)
virtualTours.POST("/property/{id}/schedule-live", accessTokenVerifierMiddleware, routes.ScheduleLiveTour)
virtualTours.GET("/property/{id}/360-images", routes.Get360Images)
virtualTours.POST("/property/{id}/360-upload", accessTokenVerifierMiddleware, routes.Upload360Images)
```

### 8. **Analytics & Reporting Module** (`/api/analytics`)
```go
analytics.GET("/property/{id}/views", accessTokenVerifierMiddleware, routes.GetPropertyAnalytics)
analytics.GET("/user/{id}/dashboard", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserDashboard)
analytics.GET("/market/trends", routes.GetMarketTrends)
analytics.GET("/landlord/{id}/performance", accessTokenVerifierMiddleware, routes.getLandlordPerformance)
```

## 🔄 Endpoints to Modify/Enhance

### Current Issues & Improvements:

#### 1. **Property Endpoints - Add Missing Features**
```go
// Add to existing property module
property.GET("/featured", routes.GetFeaturedProperties)
property.POST("/{id}/favorite", accessTokenVerifierMiddleware, routes.ToggleFavoriteProperty)
property.GET("/search/advanced", routes.AdvancedPropertySearch)
property.POST("/{id}/schedule-viewing", accessTokenVerifierMiddleware, routes.SchedulePropertyViewing)
property.GET("/{id}/similar", routes.GetSimilarProperties)
property.POST("/{id}/report", accessTokenVerifierMiddleware, routes.ReportProperty)
```

#### 2. **User Endpoints - Add Role Management**
```go
// Add to existing user module  
user.GET("/{id}/profile", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserProfile)
user.PATCH("/{id}/profile", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.UpdateUserProfile)
user.POST("/{id}/upgrade-role", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.UpgradeUserRole)
user.GET("/{id}/role", accessTokenVerifierMiddleware, utils.UserIDMiddleware, routes.GetUserRole)
```

## ❌ Endpoints to Remove/Consolidate

### 1. **Redundant Test Endpoint**
```go
// Remove this from production
notifications.POST("/test", routes.TestMessageNotification) // Keep only in development
```

### 2. **Consolidate Apartment Module**
The apartment module seems redundant. Consider merging with property:
```go
// Instead of separate apartment module, add to property:
property.GET("/{id}/units", routes.GetPropertyUnits)
property.PATCH("/{id}/units", accessTokenVerifierMiddleware, routes.UpdatePropertyUnits)
```

## 🎥 Virtual Tour Implementation Strategy

### Technical Architecture for Virtual Tours

#### 1. **Storage Strategy**
```go
// Virtual tour data structure
type VirtualTour struct {
	ID           uint      `json:"id"`
	PropertyID   uint      `json:"property_id"`
	Images360    []string  `json:"images_360"`  // S3 URLs
	VideoTour    string    `json:"video_tour"`  // S3 URL
	TourMetadata string    `json:"tour_metadata"` // JSON with hotspots, descriptions
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
}
```

#### 2. **Implementation Options**

**Option A: 360° Photo Tours**
- Upload multiple 360° photos per room
- Use libraries like Photo Sphere Viewer or Pannellum
- Store image sequences in S3
- Generate tour metadata with room connections

**Option B: Video Tours**
- Record walkthrough videos
- Store in S3 with CDN delivery
- Add interactive hotspots during playback
- Integrate with video players like Video.js

**Option C: Live Virtual Tours**
- Integration with Zoom/Google Meet APIs
- Scheduled tour booking system
- Calendar integration for landlords
- Recording capabilities for later viewing

#### 3. **Frontend Integration**
```javascript
// Example virtual tour component
const VirtualTourViewer = ({ propertyId }) => {
  const [tourData, setTourData] = useState(null);
  
  useEffect(() => {
    fetch(`/api/virtual-tours/property/${propertyId}`)
      .then(res => res.json())
      .then(setTourData);
  }, [propertyId]);

  return (
    <div className="virtual-tour-container">
      {tourData?.images_360?.map((imageUrl, index) => (
        <PhotoSphereViewer
          key={index}
          src={imageUrl}
          height="500px"
          width="100%"
        />
      ))}
    </div>
  );
};
```

## 📋 Implementation Priority Order

### Phase 1: Core Marketplace (Week 1-2)
1. Payment Module
2. Rental Applications Module
3. User Verification Module
4. Lease Management Module

### Phase 2: Trust & Safety (Week 3)
5. Enhanced User Profiles
6. Insurance Module (basic)
7. Virtual Tours (basic 360° photos)

### Phase 3: Advanced Features (Week 4-5)
8. Logistics Module
9. Analytics Module
10. Advanced Virtual Tours (live tours, video)

### Phase 4: Optimization (Week 6)
11. Performance optimization
12. Advanced search features
13. Mobile app API enhancements

## 🏗️ Database Schema Additions

You'll need these additional tables:
- `payments`
- `rental_applications` 
- `user_verifications`
- `leases`
- `insurance_policies`
- `logistics_bookings`
- `virtual_tours`
- `property_views` (analytics)

## 🔒 Security Considerations

1. **Payment Security**: PCI compliance, encrypted storage
2. **Document Verification**: Secure file upload, virus scanning
3. **Virtual Tours**: Watermarking, access control
4. **API Rate Limiting**: Implement for all new endpoints
5. **Role-Based Access**: Admin endpoints for verification approval

This roadmap will transform your existing foundation into a comprehensive rental marketplace platform while maintaining the modular monolithic architecture you're following.