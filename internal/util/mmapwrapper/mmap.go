package mmapwrapper

/*
func MRead(fileName string, offset int64, length int64) string {
	fp, _ := os.OpenFile(fileName, os.O_RDONLY, 0777)
	defer fp.Close()
	finfo, _ := fp.Stat()
	dataByte, _ := syscall.Mmap(int(fp.Fd()), offset, length, syscall.PROT_READ, syscall.MAP_SHARED)
	syscall.Munmap(dataByte)
	return tools.Bytes2Str(dataByte)
}

func MWrite(fileName string, offset int64, length int, src []byte) {
	fp, _ := os.OpenFile(fileName, os.O_WRONLY, 0777)
	defer fp.Close()
	dst, _ := syscall.Mmap(int(fp.Fd()), offset, length, syscall.PROT_WRITE, syscall.MAP_SHARED)
	copy(dst, src)
	syscall.Munmap(dst)
}*/
