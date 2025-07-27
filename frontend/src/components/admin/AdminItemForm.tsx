import React from "react";
import { useParams } from "react-router-dom";
import AdminCreateItem from "./AdminCreateItem";
import AdminEditItem from "./AdminEditItem";

interface AdminItemFormProps {
  mode: "create" | "edit";
  itemId?: string;
  onSuccess?: () => void;
  onCancel?: () => void;
}

const AdminItemForm: React.FC<AdminItemFormProps> = ({
  mode,
  itemId,
  onSuccess,
  onCancel,
}) => {
  const { id } = useParams<{ id: string }>();
  const effectiveId = itemId || id;

  if (mode === "create") {
    return <AdminCreateItem onSuccess={onSuccess} onCancel={onCancel} />;
  }

  if (mode === "edit" && effectiveId) {
    return (
      <AdminEditItem
        itemId={effectiveId}
        onSuccess={onSuccess}
        onCancel={onCancel}
      />
    );
  }

  return null;
};

export default AdminItemForm;
