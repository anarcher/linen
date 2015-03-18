package main

func Build(srcPath, targetPath string) (files Files, err error) {
	files, err = ReadFiles(srcPath)
	if err != nil {
		logger.Error("Read", "err", err)
		return
	}

	err = TransformFiles(files)
	if err != nil {
		logger.Error("Transform", "err", err)
		return
	}

	err = WriteFiles(files, targetPath)
	if err != nil {
		logger.Error("Write", "err", err)
		return
	}

	return
}
