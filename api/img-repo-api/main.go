package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	IMG_PATH string = "/images/"
)

type ImageMetadata struct {
	Width int
	Height int
	Size int64
	Format string
	Name string
}

func (this *ImageMetadata) matchesQuery(minw, maxw, minh, maxh, mins, maxs, f string) bool {
	res := false
	if minw != "" {
		minwInt, err := strconv.Atoi(minw)
		if err == nil {
			res = this.Width >= minwInt
		}
	}
	if maxw != "" {
		maxwInt, err := strconv.Atoi(maxw)
		if err == nil {
			res = this.Width <= maxwInt
		}
	}
	if minh != "" {
		minhInt, err := strconv.Atoi(minh)
		if err == nil {
			res = this.Height >= minhInt
		}
	}
	if maxh != "" {
		maxhInt, err := strconv.Atoi(maxh)
		if err == nil {
			res = this.Height <= maxhInt
		}
	}
	if mins != "" {
		minsInt, err := strconv.Atoi(mins)
		if err == nil {
			res = this.Height >= minsInt
		}
	}
	if maxs != "" {
		maxsInt, err := strconv.Atoi(maxs)
		if err == nil {
			res = this.Height <= maxsInt
		}
	}
	if f != "" {
		res = f == this.Format
	}
	return res
}

func getImageFormat(img []byte) (string, error) {
	_, format, err := image.Decode(bytes.NewBuffer(img))
	return format, err
}

func getImageMetadata(path string) (ImageMetadata, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return ImageMetadata{}, err
	} else {
		defer imgFile.Close()
		img, format, err := image.Decode(imgFile)
		if err != nil {
			return ImageMetadata{}, err
		} else {
			imgFileInfo, err := imgFile.Stat()
			imgName := strings.TrimSuffix(imgFileInfo.Name(), filepath.Ext(imgFileInfo.Name()))
			if err != nil {
				return ImageMetadata{}, err
			} else {
				return ImageMetadata{
					Width: img.Bounds().Dx(),
					Height: img.Bounds().Dy(),
					Size: imgFileInfo.Size(),
					Format: format,
					Name: imgName,
				}, nil
			}
		}
	}
}

func getImages(c *fiber.Ctx) error {
	minWidth := c.Query("minw")
	maxWidth := c.Query("maxw")
	minHeight := c.Query("minh")
	maxHeight := c.Query("maxh")
	minSize := c.Query("mins")
	maxSize := c.Query("maxs")
	format := c.Query("f")
	res := make([]ImageMetadata, 0)
	files, err := ioutil.ReadDir(IMG_PATH)
	if err != nil {
		return err
	} else {
		for _, f := range files {
			md, err := getImageMetadata(filepath.Join(IMG_PATH, f.Name()))
			if err != nil {
				return err
			} else {
				if md.matchesQuery(minWidth, maxWidth, minHeight, maxHeight, minSize, maxSize, format) {
					res = append(res, md)
				}
			}
		}
		return c.JSON(res)
	}
}

func postImage(c *fiber.Ctx) error {
	name := fmt.Sprint(time.Now().UnixNano())
	body := make([]byte, len(c.Body()))
	copy(body, c.Body())
	format, err := getImageFormat(body)
	if err != nil {
		return err
	} else {
		file, err := os.Create(filepath.Join(IMG_PATH, (name + "." + format)))
		if err != nil {
			return err
		} else {
			defer file.Close()
			_, err := file.Write(body)
			if err != nil {
				return err
			} else {
				metadata, err := getImageMetadata(filepath.Join(IMG_PATH, (name + "." + format)))
				if err != nil {
					return err
				} else {
					return c.JSON(metadata)
				}
			}
		}
	}
}

func main() {
	
	app := fiber.New()
	app.Static("/images", "/images")
	app.Get("/images/metadata", getImages)
	app.Post("/images", postImage)
	log.Fatal(app.Listen("0.0.0.0:80"))
}
