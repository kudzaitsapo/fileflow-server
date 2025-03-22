import NextAuth from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { NextAuthOptions } from "next-auth";
import axiosInstance from "@/utils/axios.unauth";

export const authOptions: NextAuthOptions = {
  session: {
    strategy: "jwt",
  },

  pages: {
    signIn: "/login",
  },
  callbacks: {
    async jwt({ token, user }) {
      if (user) {
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
        const u = user as unknown as any;
        return {
          ...token,
          accessToken: u.accessToken,
          userid: u.id,
        };
      }
      return token;
    },
    async session({ session, token }) {
      return {
        ...session,
        user: {
          ...session.user,
          id: token.userid,
          accessToken: token.accessToken,
        },
      };
    },
  },
  secret: process.env.NEXTAUTH_SECRET,
  providers: [
    CredentialsProvider({
      name: "Credentials",
      credentials: {
        email: { label: "Email", type: "email" },
        password: { label: "Password", type: "password" },
      },
      async authorize(credentials) {
        const data = {
          email: credentials?.email,
          password: credentials?.password,
        };

        try {
          const response = await axiosInstance.post("/auth/login", data);
          if (
            response.data &&
            response.data.result.user &&
            response.data.result.token
          ) {
            const { user, token } = response.data.result;
            return {
              id: user.id,
              name: user.first_name,
              email: user.email,
              accessToken: token,
            };
          }

          return null;
        } catch (e) {
          console.error("Authorization error:", e);
          return null;
        }
      },
    }),
  ],
};

export const handler = NextAuth(authOptions);

export { handler as GET, handler as POST };
