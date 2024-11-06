package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Function to compress files into a gzip archive
func compressFiles(files []string, archiveName string) (string, error) {
	// Create the compressed file
	outFile, err := os.Create(archiveName)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(outFile)
	defer gzipWriter.Close()

	// Loop through the files and add them to the gzip archive
	for _, file := range files {
		fileToCompress, err := os.Open(file)
		if err != nil {
			return "", err
		}
		defer fileToCompress.Close()

		// Copy the file content into the gzip writer
		_, err = io.Copy(gzipWriter, fileToCompress)
		if err != nil {
			return "", err
		}
	}
	return archiveName, nil
}

// Function to upload a file to AWS S3
func uploadToS3(filePath, bucketName, objectKey string) error {
	// Initialize a session with AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		return err
	}

	// Create an S3 client
	s3Client := s3.New(sess)

	// Open the file to be uploaded
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Upload the file to the specified S3 bucket
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Successfully uploaded %s to S3 bucket %s\n", objectKey, bucketName)
	return nil
}

// Function to automate the backup process
func automateBackup(files []string, bucketName, s3Path string) {
	// Compress the files into a .gzip archive
	archiveName := fmt.Sprintf("backup-%s.gz", time.Now().Format("2006-01-02_15-04-05"))
	_, err := compressFiles(files, archiveName)
	if err != nil {
		log.Fatalf("Failed to compress files: %v", err)
	}

	// Upload the archive to S3
	err = uploadToS3(archiveName, bucketName, s3Path+archiveName)
	if err != nil {
		log.Fatalf("Failed to upload to S3: %v", err)
	}

	// Optionally, delete the local archive file after upload
	err = os.Remove(archiveName)
	if err != nil {
		log.Fatalf("Failed to delete local file: %v", err)
	}

	fmt.Printf("Backup process completed successfully.\n")
}

func main() {
	// Define the files to backup (you can also dynamically list them)
	files := []string{
		"file1.txt",
		"file2.txt",
		"file3.txt",
	}

	// Define your S3 bucket and the desired path where the file will be stored
	bucketName := "your-bucket-name"
	s3Path := "backups/"

	// Start the backup automation process
	automateBackup(files, bucketName, s3Path)
}
