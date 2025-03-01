export function Footer() {
  return (
    <footer className="border-t bg-muted/50">
      <div className="container py-8 grid grid-cols-1 md:grid-cols-3 gap-8 px-2">
        <div className="space-y-2">
          <h3 className="font-semibold">ShopSphere</h3>
          <p className="text-muted-foreground text-sm">
            Your trusted online shopping destination
          </p>
        </div>
        <div className="space-y-2">
          <h4 className="font-medium">Quick links</h4>
          <ul className="space-y-1 text-sm">
            <li>
              <a href="/about" className="text-muted-foreground">
                About Us
              </a>
            </li>
            <li>
              <a href="/contact" className="text-muted-foreground">
                Contact
              </a>
            </li>
            <li>
              <a href="/faq" className="text-muted-foreground">
                FAQ
              </a>
            </li>
          </ul>
        </div>
        <div className="space-y-2">
          <h4 className="font-medium">Follow us</h4>
          <div className="flex gap-2">
            <a href="#">Facebook</a>
            <a href="#">Twitter</a>
            <a href="#">Instagram</a>
          </div>
        </div>
      </div>
    </footer>
  );
}
