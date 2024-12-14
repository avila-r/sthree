package buckets

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

// buckets.Module serves as a wrapper for the AWS S3 SDK client and provides
// methods to manage and interact with S3 buckets.
//
// Fields:
// - Sdk (*s3.S3): An instance of the AWS S3 SDK client, used to
//   communicate with the Amazon S3 service.
//
// Usage:
// The Module struct acts as a central point for bucket-related operations,
// such as listing buckets, creating new buckets, or managing existing ones.
//
// Example:
//   import (
//       "github.com/avila-r/sthree"
//   )
//
//   func main() {
//       client, _ := sthree.Session()
//
//       // List all buckets
//       buckets, err := client.Buckets.List()
//       if err != nil {
//           log.Fatalf("Failed to list buckets: %v", err)
//       }
//
//       fmt.Println("Buckets:", buckets)
//   }
type Module struct {
	// Sdk is the AWS S3 SDK client used to interact with the S3 service.
	Sdk *s3.S3
}
