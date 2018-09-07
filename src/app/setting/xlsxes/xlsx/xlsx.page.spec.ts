import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { XlsxPage } from './xlsx.page';

describe('XlsxPage', () => {
  let component: XlsxPage;
  let fixture: ComponentFixture<XlsxPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ XlsxPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(XlsxPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
