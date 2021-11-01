import Head from "next/head";
import Nav from "../components/Nav";
import React, { useState } from "react";
import { useTable } from "react-table";

export default function Home({ products }) {
  let tempProds = products;
  const [prods, setProds] = useState(tempProds);
  const data = React.useMemo(() => prods, [prods]);
  const columns = React.useMemo(
    () => [
      {
        Header: "ID",
        accessor: "id", // accessor is the "key" in the data
      },
      {
        Header: "Name",
        accessor: "name", // accessor is the "key" in the data
      },
      {
        Header: "Description",
        accessor: "description",
      },
      {
        Header: "Price",
        accessor: "price", // accessor is the "key" in the data
      },
      {
        Header: "SKU",
        accessor: "sku", // accessor is the "key" in the data
      },
    ],
    []
  );
  const tableInstance = useTable({ columns, data });
  const { getTableProps, getTableBodyProps, headerGroups, rows, prepareRow } = tableInstance;
  return (
    <div className="flex flex-col h-screen w-full items-center">
      <Head>
        <title>Coffee Products</title>
      </Head>
      <Nav prods={prods} setProds={setProds} original={products} />
      <div className="my-10 w-3/4 text-center h-full">
        <h1 className="text-5xl mb-6">Menu</h1>
        <table {...getTableProps()} className="w-full flex flex-col rounded-lg">
          <thead className="border-b border-gray-200 flex justify-between w-full rounded-lg">
            {
              // Loop over the header rows
              headerGroups.map((headerGroup) => (
                // Apply the header row props
                <tr className="flex w-full bg-gray-100 rounded-lg" {...headerGroup.getHeaderGroupProps()}>
                  {
                    // Loop over the headers in each row
                    headerGroup.headers.map((column) => (
                      // Apply the header cell props
                      <th className="flex justify-evenly w-full" {...column.getHeaderProps()}>
                        {
                          // Render the header
                          column.render("Header")
                        }
                      </th>
                    ))
                  }
                </tr>
              ))
            }
          </thead>
          {/* Apply the table body props */}
          <tbody className="flex flex-col w-fullrounded-lg" {...getTableBodyProps()}>
            {
              // Loop over the table rows
              rows.map((row) => {
                // Prepare the row for display
                prepareRow(row);
                return (
                  // Apply the row props
                  <tr className="flex bg-gray-50 w-full rounded-lg py-2" {...row.getRowProps()}>
                    {
                      // Loop over the rows cells
                      row.cells.map((cell) => {
                        // Apply the cell props
                        return (
                          <td className="flex justify-evenly w-full" {...cell.getCellProps()}>
                            {
                              // Render the cell contents
                              cell.render("Cell")
                            }
                          </td>
                        );
                      })
                    }
                  </tr>
                );
              })
            }
          </tbody>
        </table>
      </div>
    </div>
  );
}

// interface Product {
//   id,
//   name,
//   description
//   price,
//   sku
// }

export const getServerSideProps = async ({ params, res }) => {
  try {
    // const { id } = params;
    const result = await fetch(`http://localhost:9090/products`).then((response) => response.json());
    return {
      props: {
        products: result,
      },
    };
  } catch {
    res.statusCode = 404;
    return {
      props: {},
    };
  }
};
