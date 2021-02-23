import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'Feature Toggle UI';
  featureButton = 'Create New Feature';
  customerButton = 'Create New Customer';
  customerSelect = "Drop down to filter by customer"
  active = "Active Features"
  archive = "Archived Features"
  editButton = "Edit Feature"
  archiveButton = "Archive Feature"
  toggleButton = "Radio button to turn on or off"
}
