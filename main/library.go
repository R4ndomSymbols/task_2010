package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type fileName string

var (
	key   *string
	value *string
)

const path fileName = "dict.txt"

func init() {
	key = flag.String("key", "", "specified key")
	value = flag.String("value", "", "specified value")
}

func main() {

	dictionary := Read(path)
	flag.Parse()
	// добавление - указан ключ и значение
	// удаление - указан ключ но не указано значение
	// вывод на экран - ничего не указано
	if *key != "" {
		if *value != "" {
			Add_record(&dictionary, *key, *value)
			fmt.Println(fmt.Sprintf("record with key %s and value %s was added", *key, *value))
		} else {
			Delete_record(&dictionary, *key)
			fmt.Println(fmt.Sprintf("record with key %s was deleted (if existed)", *key))
		}
		Overwrite(path, &dictionary)
	} else {
		Write_all(&dictionary)
	}

}

func Read(n fileName) map[string]string {

	raw_dic, err := os.Open(string(n))

	defer raw_dic.Close()

	if err != nil {
		raw_dic, err = os.Create(string(n))
		if err != nil {
			panic("Вероятно, у приложения нет прав на создание файла или иное взаимодействие с файловой системой")
		}
	}
	fileinfo, err := raw_dic.Stat()
	filelength := fileinfo.Size()
	if filelength > 0 {
		byte_array := make([]byte, filelength)
		_, err := raw_dic.Read(byte_array)
		if err != nil {
			return make(map[string]string, 1)
		}
		str_raw_dic := string(byte_array)
		records_mas := strings.Split(str_raw_dic, "\n")
		records_mas = records_mas[:len(records_mas)-1]
		dict := make(map[string]string)
		for _, str := range records_mas {
			key_value_pair := strings.Split(str, "\t")
			if len(key_value_pair) != 2 {
				return make(map[string]string)
			}
			dict[key_value_pair[0]] = key_value_pair[1]
		}
		return dict

	} else {
		return make(map[string]string, 1)
	}
}
func Overwrite(n fileName, dict *map[string]string) {

	err := os.Remove(string(n))
	if err != nil {
		panic("Вероятно, у приложения нет прав на создание файла или иное взаимодействие с файловой системой")
	}
	raw_dic, err := os.Create(string(n))
	if err != nil {
		panic("Вероятно, у приложения нет прав на создание файла или иное взаимодействие с файловой системой")
	}

	defer raw_dic.Close()

	encoded_dict := make([]byte, 0)
	for key, value := range *dict {
		record := key + "\t" + value + "\n"
		b := []byte(record)
		encoded_dict = append(encoded_dict, b...)
	}
	_, err = raw_dic.Write(encoded_dict)
	if err != nil {
		panic(err)
	}
}

func Add_record(dictionary *map[string]string, key string, value string) {

	(*dictionary)[key] = value
}

func Delete_record(dictionary *map[string]string, key string) {

	excluded_element_dictionary := make(map[string]string)
	for ex_key, value := range *dictionary {
		if ex_key != key {
			excluded_element_dictionary[ex_key] = value
		}
	}
	*dictionary = excluded_element_dictionary
}

func Write_all(dictionary *map[string]string) {

	for key, value := range *dictionary {
		fmt.Println("key: " + key + " value : " + value)
	}
}

func Clear(n fileName) {

	err := os.Remove(string(n))
	if err != nil {
		panic("Вероятно, у приложения нет прав на создание файла или иное взаимодействие с файловой системой")
	}
	f, err := os.Create(string(n))
	if err != nil {
		panic("Вероятно, у приложения нет прав на создание файла или иное взаимодействие с файловой системой")
	}
	defer f.Close()

}
