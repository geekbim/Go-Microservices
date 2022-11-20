package exceptions

import (
	"google.golang.org/grpc/codes"
)

func MapToGrpcStatusCode(status Status) codes.Code {
	switch status {
	case ERRDOMAIN, ERRBUSSINESS:
		return codes.InvalidArgument
	case ERRNOTFOUND:
		return codes.NotFound
	case ERRAUTHORIZED:
		return codes.Unauthenticated
	case ERRFORBIDDEN:
		return codes.PermissionDenied
	case ERRSYSTEM, ERRREPOSITORY, ERRUNKNOWN:
		return codes.Internal
	default:
		return codes.Internal
	}
}
