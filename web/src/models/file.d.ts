export interface FileType {
  id: number;
  name: string;
  mimetype: string;
  description: string;
  created_at: string;
  icon: string;
}
export interface StoredFile {
  id: string;
  name: string;
  icon?: string;
  mime_type: string;
  size: number;
  uploaded_at: string;
  folder?: string;
  file_type?: FileType;
}
