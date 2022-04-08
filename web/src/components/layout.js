import Head from "next/head";
import styles from "./layout.module.css";
import Link from "next/link";
import Profile from "../components/profile";

export const siteTitle = "wonpanu";

export default function Layout({ children, home }) {
  return (
    <div className={styles.container}>
      <Head>
        <link rel="icon" href="/favicon.ico" />
        <meta name="description" content="wonpanu's personal blog" />
        <meta
          property="og:image"
          content={`https://og-image.vercel.app/${encodeURI(
            siteTitle
          )}.png?theme=light&md=0&fontSize=75px&images=https%3A%2F%2Fassets.zeit.co%2Fimage%2Fupload%2Ffront%2Fassets%2Fdesign%2Fnextjs-black-logo.svg`}
        />
        <meta name="og:title" content={siteTitle} />
        <meta name="twitter:card" content="summary_large_image" />
      </Head>
      <div className="block">
        <div className="block m-auto max-w-screen-2xl">
          <div className="flex flex-direction-row justify-between">
            <Profile />
            <main className="block pb-0 flex-grow flex-shrink min-w-0">
              <main>{children}</main>
              {!home && (
                <div className={styles.backToHome}>
                  <Link href="/">
                    <a>← Back to home</a>
                  </Link>
                </div>
              )}
            </main>
          </div>
        </div>
      </div>
    </div>
  );
}
