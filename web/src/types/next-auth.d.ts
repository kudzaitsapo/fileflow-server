import "next-auth";

declare module "next-auth" {
  interface User {
    id: number;
    email: string;
    name: string;
    accessToken: string;
  }

  interface Session extends DefaultSession {
    user: User;
    error: string;
  }

  interface JWT {
    user: User;
    access_token: string;
  }
}
