import { Fragment } from "react";
import { TbMedicalCrossFilled } from "react-icons/tb";
import { IoLogOutOutline } from "react-icons/io5";
import { useRouter } from "next/router";
import menu from "./menu";

const SideBar = ({ pathname }: sideBarProp) => {
  const router = useRouter();

  return (
    <div className="h-full w-full max-w-[200px] bg-zinc-900 flex flex-col justify-between">
      <div>
        <div
          className="py-6 flex items-center gap-4 p-6"
          onClick={() => router.push("/")}
        >
          <div className="bg-green-500 rounded-md p-1">
            <TbMedicalCrossFilled className="text-2xl" />
          </div>
          <h1>POSPharm</h1>
        </div>
        <div className="ml-6 mr-2 h-[1px] bg-zinc-800"></div>
        {Object.entries(menu).map(([key, value], main_index) => (
          <Fragment key={main_index}>
            <div className="p-6">{key}</div>
            <div className="flex flex-col gap-2">
              {value.map((item, index) => (
                <div
                  key={index}
                  onClick={() => router.push(item.endpoint)}
                  className={`flex gap-6 p-1 items-center transition duration-300 ease-in-out bg-linear-to-r/srgb ${
                    pathname === item.endpoint
                      ? "from-green-950"
                      : "hover:from-zinc-700"
                  }`}
                >
                  <div
                    className={`pl-6 py-2 text-2xl ${
                      pathname === item.endpoint ? "text-green-600" : ""
                    }`}
                  >
                    {item.icon}
                  </div>
                  {item.label}
                </div>
              ))}
            </div>
          </Fragment>
        ))}
      </div>
      <div className="flex gap-6 pb-4 items-center">
        <div className="pl-6 py-2 text-2xl">
          <IoLogOutOutline />
        </div>
        Log out
      </div>
    </div>
  );
};

export default SideBar;
