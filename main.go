package main

import (
	"fmt"
	"time"
	apb "./proto"
	)

type authzModel struct {
        ID               int64
        IdentifierType   uint8
        IdentifierValue  string
        RegistrationID   int64
        Status           uint8
        Expires          time.Time
}

func modelToAuthzPB(am *authzModel) (*apb.Authorization, error) {
        expires := am.Expires.UTC().UnixNano()
        id := fmt.Sprintf("%d", am.ID)
        status := "valid"
        pb := &apb.Authorization{
                Id:             id,
                Status:         status,
                Identifier:     am.IdentifierValue,
                RegistrationID: am.RegistrationID,
                Expires:        expires,
        }
        // snip
        return pb, nil
}


func authzModelMapToPB(m map[string]authzModel) (*apb.Authorizations, error) {
        resp := &apb.Authorizations{}
        for k, v := range m {
                // Make a copy of k because it will be reassigned with each loop.
                kCopy := k
                authzPB, err := modelToAuthzPB(&v)
                if err != nil {
                        return nil, err
                }
                resp.Authz = append(resp.Authz, &apb.Authorizations_MapElement{Domain: kCopy, Authz: authzPB})
        }
        return resp, nil
}


func main() {
	authzModels := [...]authzModel{
		authzModel{1,1,"www.example1.com",1,1, time.Date(2020, time.January, 1, 1, 1, 1, 1, time.UTC)},
		authzModel{2,2,"www.example2.com",2,2, time.Date(2020, time.February, 2, 2, 2, 2, 2, time.UTC)},
		authzModel{3,3,"www.example3.com",3,3, time.Date(2020, time.March, 3, 3, 3, 3, 3, time.UTC)},
	}
	authzModelMap := make(map[string]authzModel)
	for _, am := range authzModels {
		authzModelMap[am.IdentifierValue] = am
        }
	resp, _ := authzModelMapToPB(authzModelMap)
	fmt.Printf("%+v, Identifier:%p, RegistrationID:%p\n", resp.Authz[0], resp.Authz[0].Authz.Identifier, resp.Authz[0].Authz.RegistrationID)
	fmt.Printf("%+v, Identifier:%p, RegistrationID:%p\n", resp.Authz[1], resp.Authz[1].Authz.Identifier, resp.Authz[1].Authz.RegistrationID)
	fmt.Printf("%+v, Identifier:%p, RegistrationID:%p\n", resp.Authz[2], resp.Authz[2].Authz.Identifier, resp.Authz[2].Authz.RegistrationID)
}
