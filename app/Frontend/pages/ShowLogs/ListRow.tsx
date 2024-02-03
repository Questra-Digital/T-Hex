import { useRouter } from "next/router";
import { useState } from "react";

const ListRow = ({ _id, owner, createdat, path, duration }: any) => {
  const router = useRouter();
  const handleProject = () => {
    router.push({
      pathname: "project/" + _id,
    });
  };

  const [selectedItemId, setSelectedItemId] = useState(null);

  const handleDownload = (itemId) => {
    setSelectedItemId(itemId);
    
    window.open(`https://drive.google.com/uc?export=download&id=${itemId}`, '_blank');
  };

  return (
    <tr>
      <td className="px-5 py-5 border-b border-gray-200 bg-white text-sm">
        <div className="flex items-center">
          <div className="flex-shrink-0 w-10 h-10">
            <img
              className="w-full h-full rounded-full"
              src="https://cdn1.vectorstock.com/i/1000x1000/67/75/document-plan-icon-digital-red-vector-18086775.jpg"
              alt=""
            />
          </div>
          <div className="ml-3">
            <p className="text-gray-900 whitespace-no-wrap">{_id}</p>
          </div>
        </div>
      </td>
    
      <td className="px-5 py-5 border-b border-gray-200 bg-white text-sm">
        <p className="text-gray-900 whitespace-no-wrap">{createdat}</p>
      </td>

      <td className="px-5 py-5 border-b border-gray-200 bg-white text-sm">
        <button className="py-2 px-4 bg-blue-500 text-white rounded hover:bg-blue-600"
         onClick={() => handleDownload(path)}
        >
          Docker
        </button>
      </td>

      <td className="px-5 py-5 border-b border-gray-200 bg-white text-sm">
        <button className="py-2 px-4 bg-orange-400 text-white rounded hover:bg-orange-600"
         onClick={() => handleDownload(path)}
        >
          Selenium
        </button>
      </td>

    </tr>
  );
};

export default ListRow;
