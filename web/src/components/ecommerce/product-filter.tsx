import { Slider } from "@/components/ui/slider";
import { Label } from "../ui/label";
import { Checkbox } from "../ui/checkbox";
import { categories } from "@/constants/categories";
interface ProductFilterParams {
  selectedCategories: string[];
  onCategoryChange: (catogories: string[]) => void;
  priceRange: [number, number];
  onPriceChange: (range: [number, number]) => void;
}
export function ProductFilter({
  selectedCategories,
  onCategoryChange,
  priceRange,
  onPriceChange,
}: ProductFilterParams) {

    const handleCategoryToggle=(category:string)=>{
        const newCategories = selectedCategories.includes(category)? selectedCategories.filter(c=>c != category) : [...selectedCategories,category]
        onCategoryChange(newCategories)
    }
  return (
    <div>
      <div className="space-y-4">
        <h3 className="font-semibold">Price Range</h3>
        <Slider
          value={priceRange}
          onValueChange={(value) => onPriceChange(value as [number, number])}
          min={0}
          max={1000}
          step={10}
          minStepsBetweenThumbs={1}
        />
        <div className="flex justify-between text-sm text-muted-foreground">
          <span>${priceRange[0]}</span>
          <span>${priceRange[1]}</span>
        </div>
      </div>

      <div className="space-y-4">
        <h3 className="font-semibold">Categories</h3>
        <div className="space-y-2">
          {categories.map((category) => (
            <div key={category.slug} className="flex items-center space-x-2">
              <Checkbox
                id={category.slug}
                checked={selectedCategories.includes(category.slug)}
                onCheckedChange={()=>handleCategoryToggle(category.slug)}
              />
              <Label>{category.name}</Label>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
