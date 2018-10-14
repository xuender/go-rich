import { IonicModule } from '@ionic/angular';
import { RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { UploadPageModule } from '../setting/xlsxes/upload/upload.module';

import { CustomersPage } from './customers.page';
import { CustomerPageModule } from './customer/customer.module';
import { DownloadPageModule } from './download/download.module';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    CustomerPageModule,
    UploadPageModule,
    DownloadPageModule,
    RouterModule.forChild([{ path: '', component: CustomersPage }])
  ],
  declarations: [CustomersPage]
})
export class ContactPageModule { }
