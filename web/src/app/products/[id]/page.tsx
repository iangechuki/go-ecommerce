import ProductPage from "./productPage";
interface ProductPageProps {
  params: {
    id: string;
  };
}
export default async function Page({ params }: ProductPageProps){
    return <ProductPage id={params.id}/>
}