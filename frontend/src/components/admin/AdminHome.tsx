import React from "react";
import { Link } from "react-router-dom";

const AdminHome: React.FC = () => {
  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">管理画面ホーム</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="bg-white p-6 rounded-lg shadow">
          <h2 className="text-xl font-semibold mb-4">商品管理</h2>
          <p className="text-gray-600 mb-4">
            商品の一覧表示、編集、削除ができます。
          </p>
          <Link
            to="/admin/items"
            className="inline-block bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition-colors"
          >
            商品一覧を見る
          </Link>
        </div>
        <div className="bg-white p-6 rounded-lg shadow">
          <h2 className="text-xl font-semibold mb-4">新規商品登録</h2>
          <p className="text-gray-600 mb-4">新しい商品を登録できます。</p>
          <Link
            to="/admin/items/new"
            className="inline-block bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 transition-colors"
          >
            商品を登録する
          </Link>
        </div>
      </div>
    </div>
  );
};

export default AdminHome;
