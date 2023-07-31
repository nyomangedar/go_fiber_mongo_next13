import type { InferGetServerSidePropsType, GetServerSideProps } from "next";

type DataItem = {
    id: string;
    content: string;
    author: string;
    createDate: string; // Assuming the date is represented as a string
};

type Article = {
    status: string;
    message: string;
    data: {
        data: DataItem[];
    };
};

export default async function Page() {
    const res = await fetch("http:/127.0.0.1:8000/article");
    // http://127.0.0.1:3000/
    const data: Article = await res.json();
    return (
        <h1>
            {data.data.data.map((data) => {
                return data.content;
            })}
        </h1>
    );
}
