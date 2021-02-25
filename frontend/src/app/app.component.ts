import { Component } from '@angular/core';
import { allCustomers } from './test.customers';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent {
  title = 'Feature Toggle UI';

  customers = allCustomers
}
