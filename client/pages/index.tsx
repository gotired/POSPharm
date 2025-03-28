import MedCard from "@/components/01home/medCard";
import { useRouter } from "next/router";

import type { ReactElement } from "react";
import type { NextPageWithLayout } from "./_app";
import Layout from "@/components/layout";

const med_items = [
  "All",
  "Tablet",
  "Capsule",
  "Suppository",
  "Eyedrops",
  "Inhaler",
  "Bottle",
  "Injectable",
  "Analgesic",
  "Antibiotic",
  "Antidepressant",
  "Antifungal",
  "Antihistamine",
  "Antihypertensive",
  "Antiviral",
  "Bronchodilator",
  "Diuretic",
  "Sedative",
];

const Home: NextPageWithLayout = () => {
  const router = useRouter();
  return (
    <div className="w-full h-full p-6 bg-zinc-300 flex flex-col gap-6">
      <div className="flex justify-between">
        <h2 className="text-2xl text-black font-semibold">Medicines</h2>
        <p
          onClick={() => router.push("/supplies")}
          className="text-green-500 hover:text-green-400 underline cursor-pointer"
        >
          See all
        </p>
      </div>
      <div className="flex gap-4 max-w-[calc(100vw-250px)] overflow-x-auto scrollbar-hide">
        {med_items.map((item, index) => (
          <div
            key={index}
            className={`py-2 px-4 rounded-4xl ${
              item === "All" ? "bg-green-500" : "bg-zinc-100 text-zinc-500"
            }`}
          >
            {item}
          </div>
        ))}
      </div>
      <div className="grid grid-cols-4 gap-6">
        <MedCard
          name="Beramol"
          detail="Paracetamol"
          volume="60 ml"
          stock={12}
          price={12.5}
          unit="bottle"
        />
        <MedCard
          name="Beramol"
          detail="Paracetamol"
          volume="60 ml"
          stock={12}
          price={12.5}
          unit="bottle"
        />
        <MedCard
          name="Beramol"
          detail="Paracetamol"
          volume="60 ml"
          stock={12}
          price={12.5}
          unit="bottle"
        />
        <MedCard
          name="Beramol"
          detail="Paracetamol"
          volume="60 ml"
          stock={12}
          price={12.5}
          unit="bottle"
        />
      </div>
      <div className="flex justify-between">
        <h2 className="text-2xl text-black font-semibold">Waiting List</h2>
        <p
          onClick={() => router.push("/supplies")}
          className="text-green-500 hover:text-green-400 underline cursor-pointer"
        >
          See all
        </p>
      </div>
    </div>
  );
};

Home.getLayout = function getLayout(page: ReactElement) {
  return <Layout>{page}</Layout>;
};

export default Home;
