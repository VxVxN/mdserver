package mocks

import (
	"time"

	"github.com/VxVxN/mdserver/internal/driver/mongo/share"
	e "github.com/VxVxN/mdserver/pkg/error"
)

type MongoShare struct{}

const MockGeneratedLink = "https://vxvxn.ddns.net/share/test/6ac6d7ef-b7cd-487a-9cf5-9e2d6fe5397f"

func (ms *MongoShare) GenerateLink(username, dirName, fileName string) (string, *e.ErrObject) {
	return MockGeneratedLink, nil
}

func (ms *MongoShare) GetLinks(username string) (*share.Share, error) {
	return &share.Share{
		Owner: "test",
		ShareLinks: []share.Link{
			{
				ID:       "6ac6d7ef-b7cd-487a-9cf5-9e2d6fe5397f",
				DirName:  "test dir1",
				FileName: "test file1",
				Create:   time.Time{},
			},
			{
				ID:       "57afcf14-0b2b-4a1e-be73-d16fbd5df032",
				DirName:  "test dir1",
				FileName: "test file2",
				Create:   time.Time{},
			},
			{
				ID:       "12e446b7-ee8a-41ef-8f7c-a951968090b3",
				DirName:  "test dir2",
				FileName: "test file3",
				Create:   time.Time{},
			},
		},
	}, nil
}

func (ms *MongoShare) DeleteLink(username, link string) *e.ErrObject {
	return nil
}
