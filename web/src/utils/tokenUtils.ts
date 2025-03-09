import CryptoJS from "crypto-js";

export const getStoredToken = () => {
  if (typeof window !== "undefined") {
    const encryptedToken = localStorage.getItem("encryptedToken");
    if (encryptedToken) {
      try {
        const decryptedToken = CryptoJS.AES.decrypt(
          encryptedToken,
          process.env.NEXT_PUBLIC_ENCRYPTION_KEY || ""
        ).toString(CryptoJS.enc.Utf8);

        return decryptedToken;
      } catch (error) {
        console.error("Error decrypting token:", error);
        return null;
      }
    }
  }
  return null;
};

// const { token } = response.data;
// if (typeof window !== "undefined") {
//   const encryptedToken = CryptoJS.AES.encrypt(
//     token,
//     process.env.NEXT_PUBLIC_ENCRYPTION_KEY ||
//       "random-word-goes-here"
//   ).toString();

//   localStorage.setItem("encryptedToken", encryptedToken);
// }
