import { GraphPayments } from "./components/graph-payments";
import { TablePayments } from "./components/table-payments";

export default function Payments() {
  return (
    <div className="flex flex-col p-8 w-full">
      <header>
        <h1 className="font-montserrat text-xl text-zinc-600">Your payments</h1>
      </header>

      <div className="mt-12">
        <GraphPayments />
      </div>

      <div>
        <TablePayments />
      </div>
    </div>
  );
}
