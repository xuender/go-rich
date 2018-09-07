import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { XlsxesPage } from './xlsxes.page';

describe('XlsxesPage', () => {
  let component: XlsxesPage;
  let fixture: ComponentFixture<XlsxesPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ XlsxesPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(XlsxesPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
