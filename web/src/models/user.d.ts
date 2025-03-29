export interface Role {
  id: number;
  name: string;
  description: string;
}

export interface User {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  created_at: string;
  is_active: boolean;
  role?: Role;
}
