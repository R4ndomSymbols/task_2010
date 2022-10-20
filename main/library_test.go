package main

import (
	"testing"
)

func TestAdding(t *testing.T) {
	// добавление и считываение из файла

	Clear(path)
	dict := Read(path)
	Add_record(&dict, "0000222112121", "123123123")
	Add_record(&dict, "1nalkq[j", "lakjdfl;kajdfl;kaj;ldfkja;lksdfj")
	Add_record(&dict, "     fks;kdfk2ek", "  1ddfsdf    ")
	Add_record(&dict, "000fsdfsdf1", "12312312sdfsdf3")
	Add_record(&dict, "?000022234234asdfasdf", "123123123")
	Add_record(&dict, ",,,0000ывафыв234а2sdf", "123123123")
	Add_record(&dict, "()()(АЫАВ002342002221фывафыва12121", "фывафыва123123123")
	Add_record(&dict, "0000222112121", "12312фйцукйцукыва3123")
	Add_record(&dict, "////?????2341203000222112121", "123123123")
	Add_record(&dict, "0йцуЭЭЭЖЭЖЭЖк000222112304982121", "......12зщйук3123123")
	Overwrite(path, &dict)

	restored_dict := Read(path)

	for key, value := range dict {
		if restored_dict[key] != value {
			t.Error("mismatch")
		}
	}
	Clear(path)
}

func TestRemoving(t *testing.T) {
	// тестирование удаления

	Clear(path)
	dict := Read(path)
	Add_record(&dict, "0000222112121", "123123123")
	Add_record(&dict, "1nalkq[j", "lakjdfl;kajdfl;kaj;ldfkja;lksdfj")
	Add_record(&dict, "     fks;kdfk2ek", "  1ddfsdf    ")
	Add_record(&dict, "000fsdfsdf1", "12312312sdfsdf3")
	Add_record(&dict, "?000022234234asdfasdf", "123123123")
	Add_record(&dict, ",,,0000ывафыв234а2sdf", "123123123")
	Add_record(&dict, "()()(АЫАВ002342002221фывафыва12121", "фывафыва123123123")
	Add_record(&dict, "0000222112121", "12312фйцукйцукыва3123")
	Add_record(&dict, "////?????2341203000222112121", "123123123")
	Add_record(&dict, "0йцуЭЭЭЖЭЖЭЖк000222112304982121", "......12зщйук3123123")
	Overwrite(path, &dict)

	restored_dict := Read(path)

	Delete_record(&restored_dict, "0000222112121")
	Delete_record(&restored_dict, "     fks;kdfk2ek")
	Delete_record(&restored_dict, "000fsdfsdf1")
	Delete_record(&restored_dict, "?000022234234asdfasdf")
	Delete_record(&restored_dict, ",,,0000ывафыв234а2sdf")
	Delete_record(&restored_dict, "()()(АЫАВ002342002221фывафыва12121")
	Delete_record(&restored_dict, "0000222112121")
	Delete_record(&restored_dict, "////?????2341203000222112121")
	Delete_record(&restored_dict, "0йцуЭЭЭЖЭЖЭЖк000222112304982121")
	Delete_record(&restored_dict, "1nalkq[j")

	if len(restored_dict) != 0 {
		t.Error("record was not deleted")
	}
	Clear(path)
}
