import { IoIosNotificationsOutline } from "react-icons/io";
import { TbScan } from "react-icons/tb";
import { LuTicketPercent } from "react-icons/lu";
import { IoSearch } from "react-icons/io5";

const NavBar = () => {
  const date = new Date();

  return (
    <div className="max-h-[150px] bg-gray-100 px-6 py-8 flex justify-between gap-6 items-center">
      <div className="flex flex-col justify-between text-nowrap">
        <div className="font-bold text-2xl text-black">Drug Store</div>
        <div className="text-sm text-zinc-500">{date.toLocaleDateString()}</div>
      </div>
      <div className="rounded-xl bg-zinc-200 text-black text-sm flex flex-row items-center p-2 flex-grow">
        <div className="text-zinc-500 text-2xl flex items-center">
          <IoSearch />
        </div>
        <input
          type="text"
          className="bg-transparent outline-none flex-1 px-2"
          placeholder="Search..."
          aria-label="Search"
        />
      </div>
      <div className="text-black text-2xl outline-2 outline-zinc-300 w-fit py-1 px-2 rounded-xl flex justify-center items-center gap-2">
        <LuTicketPercent />
        <div className="text-sm text-nowrap">Use Voucher</div>
      </div>
      <div className="text-black text-2xl outline-2 outline-zinc-300 w-fit py-1 px-2 rounded-xl flex justify-center items-center gap-2">
        <TbScan />
        <div className="text-sm text-nowrap">Scan Preparation</div>
      </div>
      <div className="text-black text-2xl outline-2 outline-zinc-300 p-1 rounded-xl flex justify-center items-center">
        <IoIosNotificationsOutline />
      </div>
      <div className="h-9 w-9 bg-zinc-300 rounded-full"></div>
      <div className="flex flex-col justify-between">
        <div className="text-black text-nowrap">Mister A B</div>
        <div className="text-sm text-zinc-500">Pharmacist</div>
      </div>
    </div>
  );
};

export default NavBar;
