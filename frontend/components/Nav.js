import Link from "next/link";

export default function Nav({ prods, setProds, original, hidden }) {
  return (
    <div className="flex bg-gray-100 h-16 p-2 items-center w-full">
      <div className="text-2xl tracking-wider">Coffee Shop</div>
      <div className="text-lg pl-6 space-x-4">
        <Link className="hover:text-gray-500" href="/">
          Home
        </Link>
        <Link className="hover:text-gray-500" href="/upload">
          Upload
        </Link>
      </div>
      {!hidden && (
        <div className="ml-auto text-lg">
          <input
            className="border border-gray-300 px-2 py-1"
            placeholder="Search"
            onChange={(e) => updateList(e, prods, setProds, original)}
          />
        </div>
      )}
    </div>
  );
}

function updateList(e, prods, setProds, original) {
  const keywords = e.target.value.toLowerCase();
  prods = original;
  const updated = prods.filter(
    (prod) =>
      prod.id.toString().includes(keywords) ||
      prod.description.toLowerCase().includes(keywords) ||
      prod.name.toLowerCase().includes(keywords) ||
      prod.sku.toLowerCase().includes(keywords) ||
      prod.price.toString().includes(keywords)
  );
  setProds(updated);
}
