import { useState } from "react";

type medCardProp = {
  name: string;
  detail?: string;
  volume: string;
  stock: number;
  price: number;
  unit: string;
};

const MedCard = ({ name, detail, volume, stock, price, unit }: medCardProp) => {
  const [count, setCount] = useState<number>(0);
  return (
    <div className="rounded-xl bg-zinc-100 text-black p-4 flex flex-col gap-4">
      <div className="flex gap-2">
        <p className="text-lg font-bold">{name}</p>
        <p className="text-xs flex items-center text-zinc-500">{detail}</p>
      </div>
      <div className="grid grid-cols-2">
        <div>
          <div className="text-xs text-zinc-500">Net</div>
          <div className="text-sm font-bold">{volume}</div>
        </div>
        <div>
          <div className="text-xs text-zinc-500">Stock</div>
          <div className="text-sm font-bold">{stock.toString()} Available</div>
        </div>
      </div>
      <div className="flex gap-6">
        <div className="flex gap-1">
          <div className="flex items-start">฿</div>
          <div className="flex items-end text-xl font-bold">
            {price.toString()}
          </div>
          <div className="flex items-end">/ {unit}</div>
        </div>
        <div className="bg-zinc-300 p-1 flex justify-between flex-grow rounded-2xl">
          <div
            onClick={() => setCount((c) => c - 1)}
            className="h-6 w-6 rounded-full flex justify-center bg-zinc-100 text-green-700 font-bold cursor-pointer"
          >
            -
          </div>
          <div>{count.toString()}</div>
          <div
            onClick={() => setCount((c) => c + 1)}
            className="h-6 w-6 rounded-full flex justify-center text-zinc-100 bg-green-700 font-bold cursor-pointer"
          >
            +
          </div>
        </div>
      </div>
    </div>
  );
};

export default MedCard;
