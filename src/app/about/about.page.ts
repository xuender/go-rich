import { Component } from '@angular/core';
import { CustomerService } from '../api/customer.service';

@Component({
  selector: 'app-about',
  templateUrl: 'about.page.html',
  styleUrls: ['about.page.scss']
})
export class AboutPage {
  url: string
  constructor() {
    this.url=CustomerService.URL
  }
}
