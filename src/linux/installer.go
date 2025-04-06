package main

import (
	"io"
	"os"
	"path/filepath"
)

func inject(path string) {
	scriptDir, _ := os.Getwd()
	yarnMetaPath := filepath.Join(scriptDir, "resources", "yarnmeta")

	file, _ := os.Open(yarnMetaPath)
	defer file.Close()

	fileNameBytes, _ := io.ReadAll(file)
	fileName := string(fileNameBytes)
	src := filepath.Join(scriptDir, "resources", fileName)
	dst := filepath.Join(path, "ENA-4-DreamBBQ_Data", "StreamingAssets", "aa", "StandaloneWindows64", fileName)

	src1 := filepath.Join(scriptDir, "resources", "font_res.resS")
	dst1 := filepath.Join(path, "ENA-4-DreamBBQ_Data", "font_res.resS")

	src2 := filepath.Join(scriptDir, "resources", "resources.assets")
	dst2 := filepath.Join(path, "ENA-4-DreamBBQ_Data", "resources.assets")

	src3 := filepath.Join(scriptDir, "resources", "JoelG.ENA4.dll")
	dst3 := filepath.Join(path, "ENA-4-DreamBBQ_Data", "Managed", "JoelG.ENA4.dll")

	src4 := filepath.Join(scriptDir, "resources", "catalog.json")
	dst4 := filepath.Join(path, "ENA-4-DreamBBQ_Data", "StreamingAssets", "aa", "catalog.json")

	src5 := filepath.Join(scriptDir, "resources", "level2")
	dst5 := filepath.Join(path, "ENA-4-DreamBBQ_Data", "level2")

	src6 := filepath.Join(scriptDir, "resources", "font_modern.resS")
	dst6 := filepath.Join(path, "ENA-4-DreamBBQ_Data", "font_modern.resS")

	copyFile(src, dst)
	copyFile(src1, dst1)
	copyFile(src2, dst2)
	copyFile(src3, dst3)
	copyFile(src4, dst4)
	copyFile(src5, dst5)
	copyFile(src6, dst6)
}
func copyFile(src, dst string) error {
	srcFile, _ := os.Open(src)
	defer srcFile.Close()

	dstFile, _ := os.Create(dst)
	defer dstFile.Close()

	io.Copy(dstFile, srcFile)

	return nil
}

func installer(path string) {
	inject(path)
}
