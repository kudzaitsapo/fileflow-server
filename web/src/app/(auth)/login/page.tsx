"use client";
import React from "react";
import "./login.css";
import { ScaleLoader } from "react-spinners";
import { useRouter } from "next/navigation";
import { signIn } from "next-auth/react";
import { useSearchParams } from "next/navigation";

const Login: React.FC = () => {
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");
  const [error, setError] = React.useState<string | undefined>("");
  const [loading, setLoading] = React.useState(false);
  const router = useRouter();
  const searchParams = useSearchParams();

  const callbackUrl = searchParams.get("callbackUrl") || "/";

  const handleLogin = async () => {
    setLoading(true);

    const result = await signIn("credentials", {
      email,
      password,
      redirect: false,
    });

    if (result?.error) {
      setLoading(false);
      setError("Invalid email or password");
    } else {
      setLoading(false);
      router.push(callbackUrl);
    }
  };

  return (
    <div className="login-container">
      <div className="world-map-side">
        <div className="world-map"></div>
        <div className="user-icons">
          <div className="user-icon icon1">
            <div className="user-icon-content">👤</div>
          </div>
          <div className="user-icon icon2">
            <div className="data-icon-content"></div>
          </div>
          <div className="user-icon icon3">
            <div className="user-icon-content">👤</div>
          </div>
          <div className="user-icon icon4">
            <div className="flame-icon"></div>
          </div>
          <div className="user-icon icon5">
            <div className="user-icon-content">👤</div>
          </div>
          <div className="user-icon icon6">
            <div className="data-icon-content"></div>
          </div>
        </div>
      </div>
      <div className="login-form-side">
        <div className="logo">
          <span className="logo-icon">📁</span> FileFlow
        </div>
        <span className="entr-span">Login to access your account</span> <br />
        {error && (
          <div>
            <div className="error-container danger">{error}</div>
            <br />
          </div>
        )}
        <form>
          <div className="form-group">
            <label>
              Email <span className="danger">*</span>
            </label>
            <input
              type="text"
              onChange={(e) => setEmail(e.target.value)}
              placeholder="Enter your email"
            />
          </div>
          <div className="form-group">
            <label>
              Password <span className="danger">*</span>
            </label>
            <input
              type="password"
              id="password"
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Enter your password"
            />
          </div>
          <div className="btn-container">
            <button
              className="btn btn-primary"
              onClick={handleLogin}
              disabled={loading || !email || !password}
            >
              <ScaleLoader loading={loading} color="#fff" height={20} />
              {!loading && "Login"}
            </button>
          </div>
          <div className="forgot-password">
            <a href="#">Forgotten your password?</a>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
