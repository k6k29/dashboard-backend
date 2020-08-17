package dockerCloud

import (
	"dashboard/postgresql"
)

type Serializer struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	UseTLS    bool   `json:"use_tls"`
	TLSCaCert string `json:"tls_ca_cert"`
	TLSCert   string `json:"tls_cert"`
	TLSKey    string `json:"tls_key"`
}

func (dc *DockerCloud) Serializer() Serializer {
	var serializer Serializer
	serializer.Id = dc.BaseModel.ID
	serializer.Name = dc.Name
	serializer.Host = dc.Host
	serializer.Port = dc.Port
	serializer.UseTLS = dc.UseTLS
	serializer.TLSCaCert = dc.TLSCaCert
	serializer.TLSCert = dc.TLSCert
	serializer.TLSKey = dc.TLSKey
	return serializer
}

func ArraySerializers(dockerCloudArray []DockerCloud) []Serializer {
	var serializerArray []Serializer
	for k, _ := range dockerCloudArray {
		serializerArray = append(serializerArray, dockerCloudArray[k].Serializer())
	}
	return serializerArray
}

func (s *Serializer) Save() error {
	db := postgresql.GetInstance()
	var dockerCloudModel DockerCloud
	if s.Id != 0 {
		if querySet := db.Find(&dockerCloudModel, s.Id); querySet.Error != nil {
			return querySet.Error
		}
	}
	dockerCloudModel.Name = s.Name
	dockerCloudModel.Host = s.Host
	dockerCloudModel.Port = s.Port
	dockerCloudModel.UseTLS = s.UseTLS
	dockerCloudModel.TLSCaCert = s.TLSCaCert
	dockerCloudModel.TLSCert = s.TLSCert
	dockerCloudModel.TLSKey = s.TLSKey
	if s.Id == 0 {
		if querySet := db.Create(&dockerCloudModel); querySet.Error != nil {
			return querySet.Error
		}
	} else {
		if querySet := db.Save(&dockerCloudModel); querySet.Error != nil {
			return querySet.Error
		}
	}
	return nil
}
