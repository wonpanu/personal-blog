import Image from "next/image";
import utilStyles from "../styles/utils.module.css";

export default function profile() {
  return (
    <div className="block min-h-screen w-96 px-8">
      <div className="fixed">
        <Image
          priority
          src="/images/wonpanu.jpeg"
          className={utilStyles.borderCircle}
          height={144}
          width={144}
          alt="wonpanu"
        />
        <p className="text-xl font-extrabold mt-4">
          <span className="bg-clip-text text-transparent bg-gradient-to-r from-pink-500 to-yellow-500">
            Panumas
            <br />
            Wongsaeng
          </span>
        </p>
        <p className="text-gray-500 text-xs font-medium mt-4">
          👨‍💻 Software Engineer at WISESIGHT
        </p>
        <div className="divide-y-2 divide-zinc-100 w-56">
          <p className="text-zinc-400 text-xs mt-2">👨‍🎓 CPE30 @ KMUTT</p>
          <p className="mt-2" />
        </div>
        <p className="text-gray-500 text-xs mt-2 flex">
          <span className="font-semibold mr-1">Email:</span>
          won.panumas@gmail.com
        </p>
        <p className="text-gray-500 text-xs mt-1 flex">
          <span className="font-semibold mr-1">Linkedin:</span>
          <a
            href="https://www.linkedin.com/in/panumas-wongsaeng-135521198/"
            target="_blank"
            rel="noopener noreferrer"
            className="cursor-pointer target:"
          >
            Panumas Wongsaeng
          </a>
        </p>
        <p className="text-gray-500 text-xs mt-1 flex">
          <span className="font-semibold mr-1">GitHub:</span>
          <a
            href="https://github.com/wonpanu"
            target="_blank"
            rel="noopener noreferrer"
            className="cursor-pointer"
          >
            wonpanu
          </a>
        </p>
      </div>
    </div>
  );
}
