import { loginUser } from "@/api/auth";
import { CircularProgress } from "@mui/material";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { AiOutlineUser } from "react-icons/ai";
import { HiOutlineLockClosed } from "react-icons/hi";
import { TbMedicalCrossFilled } from "react-icons/tb";

export default function Login() {
  const router = useRouter();

  const [credentials, setCredentials] = useState({
    username: "",
    password: "",
  });
  const [loading, setLoading] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCredentials({ ...credentials, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      const data = await loginUser(credentials.username, credentials.password);
      router.push("/tenants");
    } catch (error) {
      console.log(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full flex justify-center items-center bg-zinc-300">
      <div className="bg-zinc-900 px-6 py-10 rounded-xl flex flex-col gap-10 w-100 items-center">
        <div className="flex items-center gap-4">
          <div className="bg-green-500 rounded-md p-1">
            <TbMedicalCrossFilled className="text-2xl" />
          </div>
          <h1 className="text-2xl font-bold">POSPharm</h1>
        </div>
        <form onSubmit={handleSubmit} className="flex flex-col gap-10">
          <div className="flex flex-col gap-4">
            <div className="w-75 rounded-xl bg-zinc-200 text-black text-sm flex flex-row items-center p-2">
              <div className="text-zinc-500 text-2xl flex items-center">
                <AiOutlineUser />
              </div>
              <input
                type="text"
                name="username"
                value={credentials.username}
                onChange={handleChange}
                className="bg-transparent outline-none flex-1 px-2"
                placeholder="Username or email"
                aria-label="Username"
              />
            </div>
            <div className="w-75 rounded-xl bg-zinc-200 text-black text-sm flex flex-row items-center p-2">
              <div className="text-zinc-500 text-2xl flex items-center">
                <HiOutlineLockClosed />
              </div>
              <input
                name="password"
                value={credentials.password}
                onChange={handleChange}
                type="password"
                className="bg-transparent outline-none flex-1 px-2"
                placeholder="Password"
                aria-label="Password"
              />
            </div>
          </div>
          <button
            type="submit"
            disabled={loading}
            className="w-75 rounded-xl bg-green-500 text-sm p-2 flex justify-center cursor-pointer hover:bg-green-700"
          >
            {loading ? (
              <CircularProgress size="20px" color="inherit" />
            ) : (
              "Login"
            )}
          </button>
        </form>
      </div>
    </div>
  );
}
