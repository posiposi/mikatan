package repository

import (
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var repo *itemRepository

func testMain(m *testing.M) (*gorm.DB, error) {
	// TODO GitHub ActionsでのCI/CDを考慮して、テスト用DBの接続情報を環境変数から取得するようにする
	dsn := "root:test_root_pass@tcp(test_db:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// テスト用のテーブルを作成
	err = db.AutoMigrate(&Item{})
	if err != nil {
		return nil, err
	}

	// テストデータ投入
	err = seedTestData(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func seedTestData(db *gorm.DB) error {
	items := []Item{
		{ItemID: 1, ItemName: "Item1", Description: "Description1"},
		{ItemID: 2, ItemName: "Item2", Description: "Description2"},
	}

	if err := db.Create(&items).Error; err != nil {
		return err
	}

	return nil
}

func TestGetAllItems(t *testing.T) {
	t.Run("Get All Items - Success", func(t *testing.T) {
		// Act: メソッドを実行
		items, err := repo.GetAllItems()

		// Assert: 結果を検証
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(items) != 2 {
			t.Errorf("expected 2 items, got %d", len(items))
		}

		if items[0].Name != "Item1" || items[1].Name != "Item2" {
			t.Errorf("unexpected items: %+v", items)
		}
	})
}
