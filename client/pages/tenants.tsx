import UnstyledSelectBasic from "@/components/dropdown";
import { CircularProgress } from "@mui/material";
import { useRouter } from "next/router";
import { SetStateAction, useState } from "react";
import { TbMedicalCrossFilled } from "react-icons/tb";

export default function Tenant() {
  // const router = useRouter();
  const [tenant, setTenant] = useState<string>("branch1");
  const [loading, setLoading] = useState(false);

  const handleChange = (value: string) => {
    setTenant(value);
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    try {
      console.log("fetch");
      // const data = await loginUser(credentials.username, credentials.password);
      // router.push("/tenants");
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
          <div className="flex flex-col gap-2">
            <p className="text-sm">Select Branch</p>
            <UnstyledSelectBasic
              value={tenant}
              onChange={handleChange}
              defaultValue={"branch1"}
              items={[
                { value: "branch1", label: "Branch 1" },
                { value: "branch2", label: "Branch 2" },
                { value: "branch3", label: "Branch 3" },
              ]}
            />
          </div>
          <button
            type="submit"
            disabled={loading}
            className="w-75 rounded-xl bg-green-500 text-sm p-2 flex justify-center cursor-pointer hover:bg-green-700"
          >
            {loading ? (
              <CircularProgress size="20px" color="inherit" />
            ) : (
              "Next"
            )}
          </button>
        </form>
      </div>
    </div>
  );
}
