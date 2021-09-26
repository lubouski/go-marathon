package main

import (
        "context"
        "fmt"
	"log"
	"os"
	"io"
        "io/ioutil"

        "cloud.google.com/go/storage"
        "google.golang.org/api/iterator"
)

type Cloud struct {
	bucket string
	object string
	path string
}

func (c Cloud) List() {
        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        it := storageClient.Bucket(c.bucket).Objects(ctx, nil)
	fmt.Printf("List of cloud objects at bucket %s\n", c.bucket)
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

func (c Cloud) Get() ([]byte, error) {

        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        rc, err := storageClient.Bucket(c.bucket).Object(c.object).NewReader(ctx)
        if err != nil {
                return nil, fmt.Errorf("Object(%q).NewReader: %v", c.object, err)
        }
        defer rc.Close()

        data, err := ioutil.ReadAll(rc)
        if err != nil {
                return nil, fmt.Errorf("ioutil.ReadAll: %v", err)
        }
	fmt.Printf("Data from bucket %s for object %s\n", c.bucket, c.object)
        return data, nil
}

func (c Cloud) Put() error {
        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        // Open local file.
        f, err := os.Open(c.path)
        if err != nil {
                return fmt.Errorf("os.Open: %v", err)
        }
        defer f.Close()

        // Upload an object with storage.Writer.
        wc := storageClient.Bucket(c.bucket).Object(c.object).NewWriter(ctx)
        if _, err = io.Copy(wc, f); err != nil {
                return fmt.Errorf("io.Copy: %v", err)
        }
        if err := wc.Close(); err != nil {
                return fmt.Errorf("Writer.Close: %v", err)
        }
        fmt.Printf("Object %s uploaded to bucket %s\n", c.object, c.bucket)
	fmt.Println("--------------------------------")
        return nil
}

func (c Cloud) Delete() error {

        ctx := context.Background()
        storageClient, err := storage.NewClient(ctx)
        if err != nil {
                log.Fatal(err)
        }
        defer storageClient.Close()

        o := storageClient.Bucket(c.bucket).Object(c.object)
        if err := o.Delete(ctx); err != nil {
                return fmt.Errorf("Object(%q).Delete: %v", c.object, err)
        }
	fmt.Printf("%s deleted form bucket %s\n", c.object, c.bucket)
	fmt.Println("--------------------------------")
        return nil
}

func newCloud(path, file, data string) Storage {
	return &Cloud{
		bucket: path,
		object: file,
		path: data,
	}
}
