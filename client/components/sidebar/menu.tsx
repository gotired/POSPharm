import { GrAppsRounded } from "react-icons/gr";
import { FaRegClock } from "react-icons/fa6";
import { LuChartPie } from "react-icons/lu";
import { BsPeople } from "react-icons/bs";
import { IoSettingsOutline, IoLogOutOutline } from "react-icons/io5";
import { RiUserCommunityLine } from "react-icons/ri";

const menu = {
  "Main Menu": [
    {
      label: "Home",
      icon: <GrAppsRounded />,
      endpoint: "/",
    },
    {
      label: "History",
      icon: <FaRegClock />,
      endpoint: "/history",
    },
    {
      label: "Supplies",
      icon: <LuChartPie />,
      endpoint: "/supplies",
    },
    {
      label: "Waitlist",
      icon: <BsPeople />,
      endpoint: "/waitlist",
    },
  ],
  General: [
    {
      label: "Settings",
      icon: <IoSettingsOutline />,
      endpoint: "/settings",
    },
    {
      label: "Community",
      icon: <RiUserCommunityLine />,
      endpoint: "/community",
    },
  ],
};

export default menu;
