package identitydoc

import "testing"

var testSig = `Ob3mEexQi/91fA/HMqS7L1DraJ/8T/lAblai/PrSgx6FMMPpQpi2rftc/iUcs4Uufzq0NjXkwk95
9cRES6s3T36hWgob/cutg5imhdy5++bymuzE8Z6T35pU3y3kn4eS6Yebna1atVbAFifeAqySGXCZ
l5+VTbjj/MBI7vB1cEs=`

var testDoc = `{
  "devpayProductCodes" : null,
  "privateIp" : "172.30.0.208",
  "availabilityZone" : "us-east-1a",
  "accountId" : "021124591875",
  "version" : "2010-08-31",
  "instanceId" : "i-1ddaabe5",
  "billingProducts" : null,
  "instanceType" : "t2.nano",
  "pendingTime" : "2016-09-03T15:07:45Z",
  "architecture" : "x86_64",
  "imageId" : "ami-2d39803a",
  "kernelId" : null,
  "ramdiskId" : null,
  "region" : "us-east-1"
}`

func TestDocVerification(t *testing.T) {
	doc, err := VerifyDocumentAndSignature("us-east-1", []byte(testDoc), []byte(testSig))
	if err != nil {
		t.Errorf("Error validating document %q", err)
	}
	if doc == nil {
		t.Error("Test document failed auth")
	}

	mod := testDoc + "lol"
	doc, err = VerifyDocumentAndSignature("us-east-1", []byte(mod), []byte(testSig))
	if err != ErrInvalidDocument {
		t.Error("Invalid document didn't return ErrInvalidDocument")
	}
	if doc != nil {
		t.Error("Invalid, errored document did not return nil")
	}
}
