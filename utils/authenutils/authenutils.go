package authenutils

import (
	"net/http"

	"github.com/sshindanai/bookstore-oauth-go/oauth"
	"github.com/sshindanai/bookstore-utils-go/resterrors"
)

type Authenticate struct {
	SellerID       int64
	Error          resterrors.RestErr
	waitForProcess chan bool
}

func authenticate(r *http.Request) <-chan Authenticate {
	authChannel := make(chan Authenticate)
	wait := make(chan bool)
	go func() {
		if err := oauth.AuthenticateRequest(r); err != nil {
			// TODO: return err to the caller
			authChannel <- Authenticate{waitForProcess: wait}
			<-wait
		}
		authChannel <- Authenticate{waitForProcess: wait}
		<-wait
	}()
	return authChannel
}

func callerID(r *http.Request) <-chan Authenticate {
	callerIDChannel := make(chan Authenticate)
	wait := make(chan bool)
	go func() {
		sellerID := oauth.CallerID(r)
		if sellerID == 0 {
			callerIDChannel <- Authenticate{waitForProcess: wait}
			<-wait
		}
		callerIDChannel <- Authenticate{SellerID: sellerID, waitForProcess: wait}
		<-wait
	}()
	return callerIDChannel
}

func orderChannel(auth, callerID <-chan Authenticate) <-chan Authenticate {
	ch := make(chan Authenticate)

	go func() {
		for {
			ch <- <-auth
		}
	}()

	go func() {
		for {
			ch <- <-callerID
		}
	}()

	return ch
}

func AuthenManager(r *http.Request) Authenticate {
	authChannel := orderChannel(authenticate(r), callerID(r))

	auth := <-authChannel
	callerID := <-authChannel

	auth.waitForProcess <- true
	callerID.waitForProcess <- true

	// return auth because auth has CallerID
	return auth
}
