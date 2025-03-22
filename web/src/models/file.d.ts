export interface StoredFile {
  id: string;
  name: string;
  icon?: string;
  mime_type: string;
  size: number;
  uploaded_at: string;
  folder?: string;
}
