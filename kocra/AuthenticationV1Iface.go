package kocra

import athn "k8s.io/client-go/kubernetes/typed/authentication/v1"

func CreateSelfSubjectReviewsIface(athnV1 athn.AuthenticationV1Client, namespace string) *athn.SelfSubjectReviewInterface {
	obj := athnV1.SelfSubjectReviews()
	return &obj
}

func CreateTokenReviewsIface(athnV1 athn.AuthenticationV1Client, namespace string) *athn.TokenReviewInterface {
	obj := athnV1.TokenReviews()
	return &obj
}
