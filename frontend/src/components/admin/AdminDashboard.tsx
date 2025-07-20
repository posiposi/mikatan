import React from "react";
import { Link, Outlet } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth";

const AdminDashboard: React.FC = () => {
  const { isAdmin } = useAuth();

  if (!isAdmin) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-center">
          <h1 className="text-2xl font-bold text-red-600 mb-4">
            アクセス権限がありません
          </h1>
          <p className="text-gray-600">管理者権限が必要です。</p>
        </div>
      </div>
    );
  }

  return (
    <div className="flex min-h-screen">
      <aside className="w-64 bg-gray-800 text-white">
        <div className="p-6">
          <h2 className="text-xl font-bold mb-6">管理画面</h2>
          <nav>
            <ul className="space-y-2">
              <li>
                <Link
                  to="/admin"
                  className="block p-3 rounded hover:bg-gray-700 transition-colors"
                >
                  ダッシュボード
                </Link>
              </li>
              <li>
                <Link
                  to="/admin/items"
                  className="block p-3 rounded hover:bg-gray-700 transition-colors"
                >
                  商品一覧
                </Link>
              </li>
              <li>
                <Link
                  to="/admin/items/new"
                  className="block p-3 rounded hover:bg-gray-700 transition-colors"
                >
                  商品登録
                </Link>
              </li>
            </ul>
          </nav>
        </div>
      </aside>
      <main className="flex-1 p-6">
        <Outlet />
      </main>
    </div>
  );
};

export default AdminDashboard;
