package main

import (
        "context"
        "fmt"
	"log"
	"os"
	"io"
        "io/ioutil"
	"errors"

        "cloud.google.com/go/storage"
        "google.golang.org/api/iterator"
)

type Cloud struct {
	bucket string
	object string
	path string
}

func (c *Cloud) SetPath(bucket string) error {
	for _, l := range bucket {
		switch {
		case string(l) == "*":
			return errors.New("invalid linux path provided with symbol *")
		case string(l) == "?":
			return errors.New("invalid linux pfth provided with symbol ?")
		case string(l) == "\"":
			return errors.New("invalid linux path provided with symbol \"")
		}
	}
	c.bucket = bucket
	return nil
}

func (c *Cloud) SetFile(object string) error {
        for _, l := range object {
                switch {
                case string(l) == "*":
                        return errors.New("invalid linux filename provided with symbol *")
                case string(l) == "?":
                        return errors.New("invalid linux filename provided with symbol ?")
                case string(l) == "\"":
                        return errors.New("invalid linux filename provided with symbol \"")
                case string(l) == "/":
                        return errors.New("invalid linux filename provided with symbol /")
                }
        }
        c.object = object
        return nil
}

func (c *Cloud) SetData(path string) {
	c.path = path
}

func (c Cloud) GetPath() string {
	return c.bucket
}

func (c Cloud) GetFile() string {
        return c.object
}

func (c Cloud) GetData() string {
	return c.path
}

func (c Cloud) List(bucket string) {
        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        it := storageClient.Bucket(bucket).Objects(ctx, nil)
	fmt.Printf("List of cloud objects at bucket %s\n", bucket)
	fmt.Println("--------------------------------")
        for {
                bucketAttrs, err := it.Next()
                if err == iterator.Done {
                        break
                }
                if err != nil {
                        log.Fatal(err)
                }
                fmt.Println(bucketAttrs.Name)
        }
}

func (c Cloud) Get(bucket, object string) ([]byte, error) {

        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        rc, err := storageClient.Bucket(bucket).Object(object).NewReader(ctx)
        if err != nil {
                return nil, fmt.Errorf("Object(%q).NewReader: %v", object, err)
        }
        defer rc.Close()

        data, err := ioutil.ReadAll(rc)
        if err != nil {
                return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
        }
	fmt.Printf("Data from bucket %s for object %s\n", bucket, object)
        return data, nil
}

func (c Cloud) Put(bucket, object, path string) error {
        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        // Open local file.
        f, err := os.Open(path)
        if err != nil {
                return fmt.Errorf("os.Open: %v", err)
        }
        defer f.Close()

        // Upload an object with storage.Writer.
        wc := storageClient.Bucket(bucket).Object(object).NewWriter(ctx)
        if _, err = io.Copy(wc, f); err != nil {
                return fmt.Errorf("io.Copy: %v", err)
        }
        if err := wc.Close(); err != nil {
                return fmt.Errorf("Writer.Close: %v", err)
        }
        fmt.Printf("Object %s uploaded to bucket %s\n", object, bucket)
	fmt.Println("--------------------------------")
        return nil
}

func (c Cloud) Delete(bucket, object string) error {

        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        o := storageClient.Bucket(bucket).Object(object)
        if err := o.Delete(ctx); err != nil {
                return fmt.Errorf("Object(%q).Delete: %v", object, err)
        }
	fmt.Printf("%s deleted form bucket %s\n", object, bucket)
	fmt.Println("--------------------------------")
        return nil
}

func newCloud() Storage {
	return &Cloud{
		bucket: "lubouski-golang",
		object: "golang-test-filename.txt",
		path: "/tmp/golang-test-filename.txt",
	}
}
