class Passport
  attr_reader :fields

  def initialize(input)
    @fields = input
  end

  def all_required_fields?
    all_required_fields ||= REQUIRED.all? { |field| field_present?(field) }
  end

  def valid_rules?
    return false unless all_required_fields?

    byr = get_4digits("byr")
    return false unless byr >= 1920 && byr <= 2002

    iyr = get_4digits("iyr")
    return false unless iyr >= 2010 && iyr <= 2020

    eyr = get_4digits("eyr")
    return false unless eyr >= 2020 && eyr <= 2030

    return false unless valid_hair_color?

    return false unless valid_height?

    return false unless valid_eye_color?

    return false unless valid_passport_id?

    true
  end

  private

  REQUIRED = %w(byr iyr eyr hgt hcl ecl pid).freeze

  def field_present?(field)
    Regexp.new("(#{field}:\\S+(\\s+\|$)){1}").match?(fields)
  end

  def get_4digits(field)
    Regexp.new("(#{field}:(\\d{4})(\\s+\|$)){1}").match(fields)[2].to_i
  end

  def valid_hair_color?
    Regexp.new(/(hcl:#([0-9a-f]{6})(\s+|$)){1}/).match?(fields)
  end

  def valid_height?
    field = Regexp.new(/(hgt:(\d+)(cm|in)(\s+|$)){1}/).match(fields)
    return false if field.nil?

    number = field[2].to_i
    unit = field[3]

    if unit == "cm"
      return number >= 150 && number <= 193
    elsif unit == "in"
      return number >= 59 && number <= 76
    else
      false
    end
  end

  def valid_eye_color?
    Regexp.new(/(ecl:(amb|blu|brn|gry|grn|hzl|oth)(\s+|$)){1}/).match?(fields)
  end

  def valid_passport_id?
    Regexp.new(/(pid:(\d{9})(\s+|$)){1}/).match?(fields)
  end
end

passports = File.read("input.txt").split(/\n{2,}/)

valid_fields_passports = 0
valid_rules_passports = 0

passports.each do |passport_fields|
  passport = Passport.new(passport_fields)
  valid_fields_passports+=1 if passport.all_required_fields?
  valid_rules_passports+=1 if passport.valid_rules?
end

puts valid_fields_passports
puts valid_rules_passports