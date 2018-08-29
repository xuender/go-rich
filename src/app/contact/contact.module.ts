import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { ContactPage } from './contact.page';
import { CustomerPageModule } from '../customer/customer.module';
import { UploadPageModule } from './upload/upload.module';
import { GroupPageModule } from './group/group.module';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    CustomerPageModule,
    UploadPageModule,
    GroupPageModule,
    RouterModule.forChild([{ path: '', component: ContactPage }])
  ],
  declarations: [ContactPage]
})
export class ContactPageModule {}
