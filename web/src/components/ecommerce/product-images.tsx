"use state";

import Image from "next/image";
import { useState } from "react";

interface ProductImagesProps {
  images: string[];
}
export function ProductImages({ images }: ProductImagesProps) {
  const [selectedImage, setSelectedImage] = useState(images[0]);
  return (
    <div className="space-y-4">
      <div className="relative aspect-square rounded-xl overflow-hidden">
        <Image
          src={selectedImage}
          alt="product image"
          fill
          className="object-cover"
        />
      </div>
      <div className="grid grid-cols-4 gap-2">
        {images.map((image) => (
          <button
            key={image}
            onClick={() => setSelectedImage(image)}
            className={`relative aspect-square rounded-md overflow-hidden border-2 ${
              selectedImage === image ? "border-primary" : "border-transparent"
            }`}
          >
            <Image
              src={image}
              alt="product thumbnail"
              fill
              className="object-cover"
            />
          </button>
        ))}
      </div>
    </div>
  );
}
