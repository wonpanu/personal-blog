import Head from "next/head";
import Layout, { siteTitle } from "../components/layout";
import utilStyles from "../styles/utils.module.css";
import Date from "../components/date";

import { getSortedPostsData } from "../utils/posts";

export async function getStaticProps() {
  const allPostsData = getSortedPostsData();
  return {
    props: {
      allPostsData,
    },
  };
}

export default function Home({ allPostsData }) {
  return (
    <Layout home>
      <Head>
        <title>{siteTitle}</title>
      </Head>
      <div className="max-w-2xl mx-auto pb-16 px-4 sm:pb-24 sm:px-6 lg:max-w-7xl lg:px-8">
        <div className="grid grid-cols-1 gap-y-10 gap-x-6 sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 xl:gap-x-8">
          {allPostsData.map(({ id, date, title }) => (
            <div
              key={id}
              className="bg-white rounded-lg border border-gray-200 shadow-md max-w-xs"
            >
              <a href="#">
                <img
                  className="rounded-t-lg"
                  src="/images/default-img.jpg"
                  alt=""
                />
              </a>
              <div className="p-5">
                <a href={`/posts/${id}`}>
                  <h5 className="font-bold tracking-tight text-gray-900">
                    {title}
                  </h5>
                </a>
                <small className={utilStyles.lightText}>
                  {date && <Date dateString={date} />}
                </small>
              </div>
            </div>
          ))}
        </div>
      </div>
    </Layout>
  );
}
