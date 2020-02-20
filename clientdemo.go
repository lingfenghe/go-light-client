package main

import (
	"fmt"
	"go-light-client/client"
	"math/big"
)

func main() {

	//var a = []int{1, 2, 3, 4, 5}
	//a, _ = sliceHandle.AddOrDelS(a, 1, 1)//删除切片中的下标为1的元素
	//fmt.Println(a)

	//createWeId
	//createWeId()

	//registerAuthorityIssuer
	//registerAuthorityIssuer()

	//registerCpt
	//registerCpt()

	//createCredential
	//createCredentialPojo()

	//getWeIdDocument
	//getWeIdDocument()

	//queryAuthorityIssuer
	//queryAuthorityIssuer()

	//queryCpt
	//queryCpt()

	//verifyCredentialPojo
	//verifyCredentialPojo()

}

func createWeId() {
	err := client.CreateWeId("39.106.69.186", "6001")
	fmt.Println("err =", err)
}

func registerAuthorityIssuer() {
	privateKey, _ := new(big.Int).SetString("66891022192008505617739700669391007653429796993194861158588523882789761146404", 10)
	err := client.RegisterAuthorityIssuer("39.106.69.186", "6001", "did:weid:1:0xa657855fd8b99dccb4d5684b08cebc0fbcf2e645", "WeBank10", privateKey)
	fmt.Println(err)
}

func registerCpt() {
	cptJsonSchema := `{"weid" : "Delegator WeID", "receiver": "Receiver WeID", "content": "Authorized content"}`
	weid := "did:weid:1:0xa657855fd8b99dccb4d5684b08cebc0fbcf2e645"
	cptSignature := "Dlo9glsybdLQhrhpkYU2suRfuexAaSCvoU7l2GNx4E5uQAV7uGNNDol91IfNKo7ckqfsISmwACEYhe8F279+KgE="
	cptId, cptVersion, _ := client.RegisterCpt("39.106.69.186", "6001", weid, cptJsonSchema, cptSignature)
	fmt.Println("cptId =", cptId)
	fmt.Println("cptVersion =", cptVersion)
}

func createCredentialPojo() {
	claim := `{"weid": "did:weid:1:0x5c643d55211585870b85b590c2769dadc63e2685", "receiver": "did:weid:1:0x5c643d55211585870b85b590c2769dadc63e2685", "content": "b1016358-cf72-42be-9f4b-a18fca610fca"}`
	cptId := "1032"
	expirationDate := "2021-02-17T11:48:33Z"
	issuer := "did:weid:1:0x5c643d55211585870b85b590c2769dadc63e2685"
	privateKey, _ := new(big.Int).SetString("22329131645065715219467593886919051774099926670130033608499096229769622463634", 10)
	credentialEncodeResponse, credentialJsonStr, _ := client.CreateCredentialPojo("39.106.69.186", "6001", claim, issuer, expirationDate, cptId, privateKey)
	fmt.Println("credentialJsonStr =", credentialJsonStr)

	fmt.Println("cptId =", credentialEncodeResponse.RespBody.CptId)
	fmt.Println("issuanceDate =", credentialEncodeResponse.RespBody.IssuanceDate)
	fmt.Println("context =", credentialEncodeResponse.RespBody.Context)
	fmt.Println("claim =", credentialEncodeResponse.RespBody.Claim)
	fmt.Println("claim content =", credentialEncodeResponse.RespBody.Claim["content"])
	fmt.Println("claim receiver =", credentialEncodeResponse.RespBody.Claim["receiver"])
	fmt.Println("claim weid =", credentialEncodeResponse.RespBody.Claim["weid"])
	fmt.Println("id =", credentialEncodeResponse.RespBody.Id)
	fmt.Println("proof created =", credentialEncodeResponse.RespBody.Proof.Created)
	fmt.Println("proof creator =", credentialEncodeResponse.RespBody.Proof.Creator)
	fmt.Println("proof salt =", credentialEncodeResponse.RespBody.Proof.Salt)
	fmt.Println("proof salt content =", credentialEncodeResponse.RespBody.Proof.Salt["content"])
	fmt.Println("proof salt receiver =", credentialEncodeResponse.RespBody.Proof.Salt["receiver"])
	fmt.Println("proof salt weid =", credentialEncodeResponse.RespBody.Proof.Salt["weid"])
	fmt.Println("proof signatureValue =", credentialEncodeResponse.RespBody.Proof.SignatureValue)
	fmt.Println("type =", credentialEncodeResponse.RespBody.Type[0], credentialEncodeResponse.RespBody.Type[1])
	fmt.Println("issuer =", credentialEncodeResponse.RespBody.Issuer)
	fmt.Println("expirationDate =", credentialEncodeResponse.RespBody.ExpirationDate)
}

func getWeIdDocument() {
	weid := "did:weid:1:0x76e84aeb883bba19d4f1f47327beeeafaf059c02"
	weIdDocumentInvokeResponse, err := client.GetWeIdDocument("39.106.69.186", "6001", weid)
	fmt.Println(weIdDocumentInvokeResponse, err)
	fmt.Println("id =", weIdDocumentInvokeResponse.RespBody.Id)
	fmt.Println("created =", weIdDocumentInvokeResponse.RespBody.Created)
	fmt.Println("updated =", weIdDocumentInvokeResponse.RespBody.Updated)
	fmt.Println("authentication =", weIdDocumentInvokeResponse.RespBody.Authentication)
	fmt.Println("authentication type=", weIdDocumentInvokeResponse.RespBody.Authentication[0].Type)
	fmt.Println("authentication publicKey=", weIdDocumentInvokeResponse.RespBody.Authentication[0].PublicKey)
	fmt.Println("publicKey =", weIdDocumentInvokeResponse.RespBody.PublicKey)
	fmt.Println("publicKey id =", weIdDocumentInvokeResponse.RespBody.PublicKey[0].Id)
	fmt.Println("publicKey type =", weIdDocumentInvokeResponse.RespBody.PublicKey[0].Type)
	fmt.Println("publicKey owner =", weIdDocumentInvokeResponse.RespBody.PublicKey[0].Owner)
	fmt.Println("publicKey publicKey =", weIdDocumentInvokeResponse.RespBody.PublicKey[0].PublicKey)
}

func queryAuthorityIssuer() {
	weid := "did:weid:1:0x76e84aeb883bba19d4f1f47327beeeafaf059c02"
	authorityIssuerInvokeResponse, err := client.QueryAuthorityIssuer("39.106.69.186", "6001", weid)
	fmt.Println(authorityIssuerInvokeResponse, err)
	fmt.Println("created =", authorityIssuerInvokeResponse.RespBody.Created)
	fmt.Println("accValue =", authorityIssuerInvokeResponse.RespBody.AccValue)
	fmt.Println("name =", authorityIssuerInvokeResponse.RespBody.Name)
	fmt.Println("weid =", authorityIssuerInvokeResponse.RespBody.WeId)
}

func queryCpt() {
	cptId := "1029"
	cptInvokeResponse, err := client.QueryCpt("39.106.69.186", "6001", cptId)
	fmt.Println(cptInvokeResponse, err)
	fmt.Println("cptBaseInfo cptId =", cptInvokeResponse.RespBody.CptBaseInfo.CptId)
	fmt.Println("cptBaseInfo cptVersion =", cptInvokeResponse.RespBody.CptBaseInfo.CptVersion)
	fmt.Println("cptJsonSchema =", cptInvokeResponse.RespBody.CptJsonSchema)
	fmt.Println("cptJsonSchema schema =", cptInvokeResponse.RespBody.CptJsonSchema["$schema"])
	fmt.Println("cptJsonSchema content =", cptInvokeResponse.RespBody.CptJsonSchema["content"])
	fmt.Println("cptJsonSchema receiver =", cptInvokeResponse.RespBody.CptJsonSchema["receiver"])
	fmt.Println("cptJsonSchema type =", cptInvokeResponse.RespBody.CptJsonSchema["type"])
	fmt.Println("cptJsonSchema weid =", cptInvokeResponse.RespBody.CptJsonSchema["weid"])
	fmt.Println("metaData cptPublisher =", cptInvokeResponse.RespBody.MetaData.CptPublisher)
	fmt.Println("metaData cptSignature =", cptInvokeResponse.RespBody.MetaData.CptSignature)
	fmt.Println("metaData created =", cptInvokeResponse.RespBody.MetaData.Created)
	fmt.Println("metaData updated =", cptInvokeResponse.RespBody.MetaData.Updated)
}

func verifyCredentialPojo() {
	credentialJsonStr := `{
							"cptId": 1032,
							"issuanceDate": 1582192501,
							"context": "https://github.com/WeBankFinTech/WeIdentity/blob/master/context/v1",
							"claim": {
								"content": "b1016358-cf72-42be-9f4b-a18fca610fca",
								"receiver": "did:weid:1:0x5c643d55211585870b85b590c2769dadc63e2685",
								"weid": "did:weid:1:0x5c643d55211585870b85b590c2769dadc63e2685"
							},
							"id": "512a3c44-5376-41f1-9ab7-9d9c004f8a2f",
							"proof": {
								"created": 1582192501,
								"creator": "did:weid:1:0x5c643d55211585870b85b590c2769dadc63e2685#keys-0",
								"salt": {
									"content": "w2yzU",
									"receiver": "HazXN",
									"weid": "2iNvM"
								},
								"signatureValue": "lZbymKlIpDZid56NpxanbOpMBS/1ur5gQlnwm2CnnaYrsOT6q9lsgmpoXD8/azvtgpygPJ1UyTutjswQ29pFEQE=",
								"type": "Secp256k1"
							},
							"type": [
								"VerifiableCredential",
								"hashTree"
							],
							"issuer": "did:weid:1:0x5c643d55211585870b85b590c2769dadc63e2685",
							"expirationDate": 1613533713
						}`
	verifyCredentialInvokeResponse, err := client.VerifyCredentialPojo("39.106.69.186", "6001", credentialJsonStr)
	fmt.Println(verifyCredentialInvokeResponse)
	fmt.Println("err =", err)
}

