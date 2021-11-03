import React, { useState } from "react";
import { useRouter } from "next/router";
import Head from "next/head";
import Nav from "../components/Nav";
import axios from "axios";

function handleSubmit(e, id, file, r) {
  e.preventDefault();

  const data = new FormData();
  data.append("file", file);
  data.append("id", id);
  console.log(id, file);

  // upload the file
  axios
    .post("http://localhost:9091", data, { "content-type": `multipart/form-data; boundary=${data._boundary}` })
    .then((res) => {
      console.log(res);

      if (res.status === 200) {
        alert("Successful upload!");
        r.push("/");
      } else {
        alert("Failed upload!");
        r.push("/upload");
      }
    });
}

export default function Upload() {
  const [file, setFile] = useState(null);
  const [id, setId] = useState(null);
  const router = useRouter();

  return (
    <div className="flex flex-col h-screen w-full items-center">
      <Head>
        <title>Coffee Products</title>
      </Head>
      <Nav hidden={true} />
      <section className="h-full w-full flex justify-start items-center flex-col py-16 px-8">
        <div className="w-3/4">
          <h1 className="text-5xl mb-6">Upload</h1>
          <form
            onSubmit={(e) => handleSubmit(e, id, file, router)}
            className="flex flex-col space-y-6 w-76 items-start"
          >
            <div className="flex items-start space-x-6">
              <label className="mt-2">Product ID:</label>
              <div className="flex flex-col space-y-1">
                <input
                  onChange={(e) => setId(e.target.value)}
                  type="number"
                  min="0"
                  className="   px-2 w-16 text-center
      py-2
      bg-white
      rounded-md
      shadow-md
      tracking-wide
      uppercase
      border border-blue
      cursor-pointer
      hover:bg-purple-600 hover:text-white
      text-purple-600
      ease-linear
      transition-all
      duration-150"
                />
                <span className="text-gray-500 text-sm">Enter a product ID to upload the image</span>
              </div>
            </div>
            <div className="flex items-center space-x-20">
              <label>File:</label>
              {file && (
                <div className="inline-flex space-x-2 items-center">
                  <div>{file.name}</div>
                  <div
                    onClick={() => setFile(null)}
                    className="
                    px-3
                    py-1
                    rounded-full
                    bg-white
                    rounded-md
                    shadow-md
                    tracking-wide
                    border border-blue
                    cursor-pointer
                    hover:bg-purple-600 hover:text-white
                    text-purple-600
                    ease-linear
                    transition-all
                    duration-150"
                  >
                    x
                  </div>
                </div>
              )}
              {!file && (
                <div className="flex flex-col space-y-1 items-start">
                  <label
                    className="w-40
                      flex flex-col
                      items-center
                      px-2
                      py-2
                      bg-white
                      rounded-md
                      shadow-md
                      tracking-wide
                      uppercase
                      border border-blue
                      cursor-pointer
                      hover:bg-purple-600 hover:text-white
                      text-purple-600
                      ease-linear
                      transition-all
                      duration-150
                    "
                  >
                    <span className="text-base leading-normal">Select a file</span>
                    <input type="file" className="hidden" onChange={(e) => setFile(e.target.files[0])} />
                  </label>

                  <span className="text-gray-500 text-sm">Image to associate with product</span>
                </div>
              )}
            </div>
            <button
              type="submit"
              className="ml-28      px-2
      py-2
      bg-white
      rounded-md
      shadow-md
      tracking-wide
      uppercase
      border border-blue
      cursor-pointer
      hover:bg-purple-600 hover:text-white
      text-purple-600
      ease-linear
      transition-all
      duration-150"
            >
              Submit
            </button>
          </form>
        </div>
      </section>
    </div>
  );
}
